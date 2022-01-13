package server

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func (s *CourseAPIServer) GetLesson(ctx context.Context, req *api.GetLessonRequest) (*api.Lesson, error) {
	lesson := models.Lesson{}
	result := s.Db.First(&lesson, "id = ?", req.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &api.Lesson{
		ID:          lesson.ID.String(),
		UserId:      lesson.UserID,
		CourseId:    lesson.CourseID,
		Name:        lesson.CourseID,
		Description: lesson.Description,
		Slug:        lesson.Slug,
	}, nil
}
