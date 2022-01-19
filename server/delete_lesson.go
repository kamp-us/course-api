package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CourseAPIServer) DeleteLesson(ctx context.Context, req *api.DeleteLessonRequest) (*emptypb.Empty, error) {
	if err := validateDeleteLessonRequest(req); err != nil {
		return nil, err
	}
	err := s.backend.DeleteLesson(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteLessonRequest(req *api.DeleteLessonRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
