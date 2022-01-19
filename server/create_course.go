package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (*api.Course, error) {
	if err := validateCreateCourseRequest(req); err != nil {
		return nil, err
	}

	course, err := s.backend.CreateCourse(ctx, req.UserId, req.Name, req.Description, req.CategoryIds)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return helper.ConvertToCourseModel(course), nil
}

func validateCreateCourseRequest(req *api.CreateCourseRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	if req.UserId == "" {
		return twirp.RequiredArgumentError("user_id")
	}
	if req.Description == "" {
		return twirp.RequiredArgumentError("description")
	}
	if len(req.CategoryIds) == 0 {
		return twirp.RequiredArgumentError("category_ids")
	}
	return nil
}
