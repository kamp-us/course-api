package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) CreateLesson(ctx context.Context, req *api.CreateLessonRequest) (*api.Lesson, error) {
	if err := validateCreateLessonRequest(req); err != nil {
		return nil, err
	}

	lesson, err := s.backend.CreateLesson(ctx, req.UserId, req.Name, req.Description, req.CourseId)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.Lesson{
		ID:          lesson.ID.String(),
		UserId:      lesson.UserID,
		CourseId:    lesson.CourseID,
		Name:        lesson.Name,
		Description: lesson.Description,
		Slug:        lesson.Slug,
		//VideoId:     lesson.VideoID,
	}, nil
}

func validateCreateLessonRequest(req *api.CreateLessonRequest) error {
	if req.Name == "" {
		return twirp.RequiredArgumentError("name")
	}
	if req.UserId == "" {
		return twirp.RequiredArgumentError("user_id")
	}
	if req.Description == "" {
		return twirp.RequiredArgumentError("description")
	}
	if req.CourseId == "" {
		return twirp.RequiredArgumentError("course_id")
	}
	//if req.VideoId == "" {
	//	return twirp.RequiredArgumentError("video_id")
	//}
	return nil
}
