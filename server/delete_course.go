package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *CourseAPIServer) DeleteCourse(ctx context.Context, req *api.DeleteCourseRequest) (*emptypb.Empty, error) {
	if err := validateDeleteCourseRequest(req); err != nil {
		return nil, err
	}
	err := s.backend.DeleteCourse(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &emptypb.Empty{}, nil
}

func convertToStringPtr(value *wrapperspb.StringValue) *string {
	val := value.GetValue()
	return &val
}

func validateDeleteCourseRequest(req *api.DeleteCourseRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
