package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) GetLesson(ctx context.Context, req *api.GetLessonRequest) (*api.Lesson, error) {
	if err := validateGetLessonRequest(req); err != nil {
		return nil, err
	}

	lesson, err := s.backend.GetLesson(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return helper.ConvertToLessonModel(lesson), nil
}

func validateGetLessonRequest(req *api.GetLessonRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
