package main

import (
	"context"
	"errors"
	"net/http"
	"sync/atomic"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/validate"
	"github.com/FlowSeer/fail"
	"github.com/FlowSeer/service"
	"github.com/lewdzifer/pidx/proto/generated/go/pidx/tag/v1/tagv1connect"
	"github.com/lewdzifer/pidx/service/tag/api"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	ServiceName    = "tag-api"
	ServiceDomain  = "pidx.io/tag"
	ServiceVersion = "0.0.1"
)

type Service struct {
	running  atomic.Bool
	shutdown atomic.Bool

	server *http.Server
}

func (s *Service) Name() string {
	return ServiceName
}

func (s *Service) Namespace() string {
	return ServiceDomain
}

func (s *Service) Version() string {
	return ServiceVersion
}

func (s *Service) Health() service.Health {
	status := service.HealthStatusUnknown
	if s.running.Load() {
		status = service.HealthStatusHealthy
	}
	if s.shutdown.Load() {
		status = service.HealthStatusShutdown
	}

	return service.Health{
		Status: status,
	}
}

func (s *Service) Initialize(ctx *service.Context) error {
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return fail.WrapC(ctx, err, "failed to create validation interceptor")
	}

	otelInterceptor, err := otelconnect.NewInterceptor(
		otelconnect.WithMeterProvider(service.MeterProvider(ctx)),
		otelconnect.WithTracerProvider(service.TracerProvider(ctx)),
	)
	if err != nil {
		return fail.WrapC(ctx, err, "failed to create otel interceptor")
	}

	reflector := grpcreflect.NewStaticReflector(
		tagv1connect.TagServiceName,
	)

	mux := http.NewServeMux()

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	mux.Handle(grpchealth.NewHandler(s))
	mux.Handle(tagv1connect.NewTagServiceHandler(&api.Server{},
		connect.WithInterceptors(
			otelInterceptor,
			validateInterceptor,
		),
	))

	s.server = &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	return nil
}

func (s *Service) Run(ctx *service.Context) error {
	if s.running.Swap(true) {
		return fail.MsgC(ctx, "already running")
	}

	ctx.Info("serving API", "addr", s.server.Addr)

	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return fail.WrapC(ctx, err, "failed to start server")
	}
}

func (s *Service) Shutdown(ctx *service.Context) error {
	if !s.running.Load() {
		return nil
	}
	if s.shutdown.Swap(true) {
		return fail.MsgC(ctx, "already shut down")
	}

	return fail.WrapC(ctx, s.server.Shutdown(ctx), "failed to shutdown server")
}

func (s *Service) Check(ctx context.Context, r *grpchealth.CheckRequest) (*grpchealth.CheckResponse, error) {
	var status grpchealth.Status
	switch s.Health().Status {
	case service.HealthStatusHealthy:
		status = grpchealth.StatusServing
	default:
		status = grpchealth.StatusNotServing
	}

	return &grpchealth.CheckResponse{
		Status: status,
	}, nil
}
