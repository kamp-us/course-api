package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetCourse(ctx context.Context, id string) (*models.Course, error) {
	course := models.Course{}
	result := b.DB.First(&course, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	query := b.DB.Model(&course).Association("Categories")
	if query.Error != nil {
		return nil, query.Error
	}

	return &course, nil
}
