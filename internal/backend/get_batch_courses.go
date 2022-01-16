package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetBatchCourses(ctx context.Context, ids []string) ([]*models.Course, error) {
	var courses []*models.Course
	result := b.DB.Find(&courses, ids)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, course := range courses {
		query := b.DB.Preload("Categories").Find(&course)
		if query.Error != nil {
			return nil, query.Error
		}
	}

	return courses, nil
}
