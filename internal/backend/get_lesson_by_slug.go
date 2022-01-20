package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetLessonBySlug(ctx context.Context, slug string) (*models.Lesson, error) {
	lesson := models.Lesson{}
	result := b.DB.First(&lesson, "slug = ?", slug)
	if result.Error != nil {
		return nil, result.Error
	}

	query := b.DB.Preload("Categories").Find(&lesson)
	if query.Error != nil {
		return nil, query.Error
	}

	return &lesson, nil
}
