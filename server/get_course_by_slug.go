package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) GetCourseBySlug(ctx context.Context, req *api.GetCourseBySlugRequest) (*api.GetCourseBySlugResponse, error) {
	if err := validateGetCourseBySlugRequest(req); err != nil {
		return nil, err
	}

	course, err := s.backend.GetCourseBySlug(ctx, req.Slug)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.GetCourseBySlugResponse{Course: course.ToAPIModel()}, nil
}

func validateGetCourseBySlugRequest(req *api.GetCourseBySlugRequest) error {
	if req.Slug == "" {
		return twirp.RequiredArgumentError("slug")
	}
	return nil
}
