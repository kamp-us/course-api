package server

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func (s *CourseAPIServer) CreateLesson(ctx context.Context, req *api.CreateLessonRequest) (*api.Lesson, error) {
	lesson := models.Lesson{
		Slug:        slug.MakeLang(req.Name, "tr"),
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserId,
		CourseID:    req.CourseId,
		//VideoID:     req.VideoId,
	}

	result := s.Db.Create(&lesson)
	if result.Error != nil {
		return nil, result.Error
	}

	return &api.Lesson{
		ID:       lesson.ID.String(),
		UserId:   lesson.UserID,
		CourseId: lesson.CourseID,
		//VideoId:     lesson.VideoID,
		Name:        lesson.Name,
		Description: lesson.Description,
		Slug:        lesson.Slug,
	}, nil
}
