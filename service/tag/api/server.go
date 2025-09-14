package api

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/lewdzifer/pidx/proto/generated/go/pidx/tag/v1"
)

type Server struct{}

func (s *Server) CreateTag(ctx context.Context, c *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.CreateTagResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetTag(ctx context.Context, c *connect.Request[v1.GetTagRequest]) (*connect.Response[v1.GetTagResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateTag(ctx context.Context, c *connect.Request[v1.UpdateTagRequest]) (*connect.Response[v1.UpdateTagResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetTagAssignments(ctx context.Context, c *connect.Request[v1.GetTagAssignmentsRequest]) (*connect.Response[v1.GetTagAssignmentsResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateTagAssignment(ctx context.Context, c *connect.Request[v1.CreateTagAssignmentRequest]) (*connect.Response[v1.CreateTagAssignmentResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteTagAssignment(ctx context.Context, c *connect.Request[v1.DeleteTagAssignmentRequest]) (*connect.Response[v1.DeleteTagAssignmentResponse], error) {
	//TODO implement me
	panic("implement me")
}
