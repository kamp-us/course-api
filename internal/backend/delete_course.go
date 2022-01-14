package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) DeleteCourse(ctx context.Context, id string) error {
	course := models.Course{}
	result := b.DB.First(&course, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = b.DB.Delete(&course)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
