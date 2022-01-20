package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) GetLessonBySlug(ctx context.Context, req *api.GetLessonBySlugRequest) (*api.GetLessonBySlugResponse, error) {
	if err := validateGetLessonBySlugRequest(req); err != nil {
		return nil, err
	}

	lesson, err := s.backend.GetLessonBySlug(ctx, req.Slug)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.GetLessonBySlugResponse{Lesson: helper.ConvertToLessonModel(lesson)}, nil
}

func validateGetLessonBySlugRequest(req *api.GetLessonBySlugRequest) error {
	if req.Slug == "" {
		return twirp.RequiredArgumentError("slug")
	}
	return nil
}
