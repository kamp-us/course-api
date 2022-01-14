package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) CreateLesson(ctx context.Context, userID string, name string, description string, courseID string) (*models.Lesson, error) {
	lesson := models.Lesson{
		Slug:        slug.MakeLang(name, "tr"),
		Name:        name,
		Description: description,
		UserID:      userID,
		CourseID:    courseID,
	}

	result := b.DB.Create(&lesson)
	if result.Error != nil {
		return nil, result.Error
	}

	return &lesson, nil
}
