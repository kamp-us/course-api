package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetBatchLessons(ctx context.Context, ids []string) ([]*models.Lesson, error) {
	var lessons []*models.Lesson
	result := b.DB.Find(&lessons, ids)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, lesson := range lessons {
		query := b.DB.Preload("Categories").Find(&lesson)
		if query.Error != nil {
			return nil, query.Error
		}
	}

	return lessons, nil
}
