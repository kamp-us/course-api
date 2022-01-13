package server

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func (s *CourseAPIServer) UpdateLesson(ctx context.Context, req *api.UpdateLessonRequest) (*api.Lesson, error) {
	lesson := models.Lesson{}
	result := s.Db.First(&lesson, "id = ?", req.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	updates := models.Lesson{}

	if name := convertToStringPtr(req.Name); name != nil {
		updates.Name = *name
		updates.Slug = slug.MakeLang(*name, "tr")
	}
	if description := convertToStringPtr(req.Description); description != nil {
		updates.Description = *description
	}

	result = s.Db.Model(&lesson).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return &api.Lesson{
		ID:          lesson.ID.String(),
		UserId:      lesson.UserID,
		Name:        lesson.Name,
		Description: lesson.Description,
		Slug:        lesson.Slug,
	}, nil
}
