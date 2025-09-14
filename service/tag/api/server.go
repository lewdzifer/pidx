package api

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"connectrpc.com/connect"
	"github.com/FlowSeer/fail"
	"github.com/FlowSeer/service"
	"github.com/jackc/pgx/v5/pgxpool"
	v1 "github.com/lewdzifer/pidx/proto/generated/go/pidx/tag/v1"
)

type Handler struct {
	pool *pgxpool.Pool
}

func NewHandlerFromEnv(ctx context.Context) (*Handler, error) {
	pgHostVar := service.EnvName(service.Name(ctx), "PG_HOST")
	pgPortVar := service.EnvName(service.Name(ctx), "PG_PORT")
	pgUserVar := service.EnvName(service.Name(ctx), "PG_USER")
	pgPassVar := service.EnvName(service.Name(ctx), "PG_PASS")
	pgDbVar := service.EnvName(service.Name(ctx), "PG_DB")
	pgSslmodeVar := service.EnvName(service.Name(ctx), "PG_SSLMODE")

	pgHost, ok := os.LookupEnv(pgHostVar)
	if !ok {
		return nil, fail.New().
			Attribute("var", pgHostVar).
			Context(ctx).
			Msgf("env var must be set: %s", pgHostVar)
	}

	pgPort, ok := os.LookupEnv(pgPortVar)
	if !ok {
		pgPort = "5432"
	}

	pgPortInt, err := strconv.Atoi(pgPort)
	if err != nil {
		return nil, fail.New().
			Attribute("var", pgPortVar).
			Context(ctx).
			Msg("value must be an integer")
	}

	pgUser, ok := os.LookupEnv(pgUserVar)
	if !ok {
		pgUser = "postgres"
	}

	pgPass, ok := os.LookupEnv(pgPassVar)
	if !ok {
		pgPass = ""
	}

	pgDb, ok := os.LookupEnv(pgDbVar)
	if !ok {
		pgDb = "postgres"
	}

	pgSslmode, ok := os.LookupEnv(pgSslmodeVar)
	if !ok {
		pgSslmode = "disable"
	}

	pool, err := pgxpool.New(ctx,
		fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", pgUser, pgPass, pgHost, pgPortInt, pgDb, pgSslmode),
	)
	if err != nil {
		return nil, fail.WrapC(ctx, err, "failed to create postgres pool")
	}

	return NewHandler(ctx, pool)
}

func NewHandler(ctx context.Context, pool *pgxpool.Pool) (*Handler, error) {
	if err := pool.Ping(ctx); err != nil {
		return nil, fail.WrapC(ctx, err, "ping failed")
	}

	return &Handler{
		pool: pool,
	}, nil
}

func (s *Handler) CreateTag(ctx context.Context, c *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.CreateTagResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Handler) GetTag(ctx context.Context, c *connect.Request[v1.GetTagRequest]) (*connect.Response[v1.GetTagResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Handler) UpdateTag(ctx context.Context, c *connect.Request[v1.UpdateTagRequest]) (*connect.Response[v1.UpdateTagResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Handler) GetTagAssignments(ctx context.Context, c *connect.Request[v1.GetTagAssignmentsRequest]) (*connect.Response[v1.GetTagAssignmentsResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Handler) CreateTagAssignment(ctx context.Context, c *connect.Request[v1.CreateTagAssignmentRequest]) (*connect.Response[v1.CreateTagAssignmentResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Handler) DeleteTagAssignment(ctx context.Context, c *connect.Request[v1.DeleteTagAssignmentRequest]) (*connect.Response[v1.DeleteTagAssignmentResponse], error) {
	//TODO implement me
	panic("implement me")
}
