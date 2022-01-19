package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s CourseAPIServer) GetBatchCourses(ctx context.Context, req *api.GetBatchCoursesRequest) (*api.GetBatchCoursesResponse, error) {
	if err := validateGetBatchCoursesRequest(req); err != nil {
		return nil, err
	}

	courses, err := s.backend.GetBatchCourses(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	var batch []*api.Course
	for _, model := range courses {
		course := helper.ConvertToCourseModel(model)
		batch = append(batch, course)
	}

	return &api.GetBatchCoursesResponse{Courses: batch}, nil
}

func validateGetBatchCoursesRequest(req *api.GetBatchCoursesRequest) error {
	if len(req.Ids) == 0 {
		return twirp.RequiredArgumentError("ids")
	}
	return nil
}
