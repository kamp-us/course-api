package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server/helper"
	"github.com/twitchtv/twirp"
)

func (s CourseAPIServer) GetBatchLessons(ctx context.Context, req *api.GetBatchLessonsRequest) (*api.GetBatchLessonsResponse, error) {
	if err := validateGetBatchLessonsRequest(req); err != nil {
		return nil, err
	}
	lessons, err := s.backend.GetBatchLessons(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	var batch []*api.Lesson
	for _, model := range lessons {
		lesson := helper.ConvertToLessonModel(model)
		batch = append(batch, lesson)
	}

	return &api.GetBatchLessonsResponse{Lessons: batch}, nil
}

func validateGetBatchLessonsRequest(req *api.GetBatchLessonsRequest) error {
	if len(req.Ids) == 0 {
		return twirp.RequiredArgumentError("ids")
	}
	return nil
}
