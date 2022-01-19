package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetCoursesByCategoryID(ctx context.Context, id string) ([]*models.Course, error) {
	var courses []*models.Course
	var courseCategories []*models.CourseCategory

	query := b.DB.Preload("Course.Categories").Preload("Course").Find(&courseCategories, "category_id = ?", id)
	if query.Error != nil {
		return nil, query.Error
	}

	for _, model := range courseCategories {
		courses = append(courses, &model.Course)
	}

	return courses, nil
}
