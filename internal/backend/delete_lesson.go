package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) DeleteLesson(ctx context.Context, id string) error {
	lesson := models.Lesson{}
	if query := b.DB.First(&lesson, "id = ?", id); query.Error != nil {
		return query.Error
	}
	result := b.DB.Delete(&lesson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
