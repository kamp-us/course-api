package server

import (
	"context"

	api "github.com/kamp-us/course-api/rpc/course-api"

	"github.com/twitchtv/twirp"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CourseAPIServer) UpdateCourse(ctx context.Context, req *api.UpdateCourseRequest) (*emptypb.Empty, error) {

	if err := validateUpdateCourseRequest(req); err != nil {
		return nil, err
	}

	err := s.backend.UpdateCourse(ctx, req.ID, convertToStringPtr(req.Name), convertToStringPtr(req.Description))
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &emptypb.Empty{}, nil
}

func validateUpdateCourseRequest(req *api.UpdateCourseRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("id")
	}
	return nil
}
