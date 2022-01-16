package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetLesson(ctx context.Context, id string) (*models.Lesson, error) {
	lesson := models.Lesson{}

	result := b.DB.First(&lesson, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	query := b.DB.Preload("Categories").Find(&lesson)
	if query.Error != nil {
		return nil, query.Error
	}

	return &lesson, nil
}
