package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetCourseBySlug(ctx context.Context, slug string) (*models.Course, error) {
	course := models.Course{}
	result := b.DB.First(&course, "slug = ?", slug)
	if result.Error != nil {
		return nil, result.Error
	}

	query := b.DB.Preload("Categories").Find(&course)
	if query.Error != nil {
		return nil, query.Error
	}

	return &course, nil
}
