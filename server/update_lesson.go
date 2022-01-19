package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CourseAPIServer) UpdateLesson(ctx context.Context, req *api.UpdateLessonRequest) (*emptypb.Empty, error) {
	if err := validateUpdateLessonRequest(req); err != nil {
		return nil, err
	}

	if err := s.backend.UpdateLesson(ctx, req.ID,
		convertToStringPtr(req.Name),
		convertToStringPtr(req.Description)); err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &emptypb.Empty{}, nil
}

func validateUpdateLessonRequest(req *api.UpdateLessonRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
