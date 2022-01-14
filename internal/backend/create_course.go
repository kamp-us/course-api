package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) CreateCourse(ctx context.Context, userID string, name string, description string) (*models.Course, error) {
	course := models.Course{
		Slug:        slug.MakeLang(name, "tr"),
		Name:        name,
		Description: description,
		UserID:      userID,
	}

	result := b.DB.Create(&course)
	if result.Error != nil {
		return nil, result.Error
	}

	return &course, nil
}
