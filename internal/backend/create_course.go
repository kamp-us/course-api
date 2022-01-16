package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) CreateCourse(ctx context.Context, userID string, name string, description string, categoryIDs []string) (*models.Course, error) {

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

	var categories []*models.CourseCategory
	for _, categoryId := range categoryIDs {
		courseCategory := models.CourseCategory{CourseID: course.ID, CategoryID: categoryId}
		categories = append(categories, &courseCategory)
	}

	err := b.DB.Model(&course).Association("Categories").Append(categories)
	if err != nil {
		return nil, err
	}

	return &course, nil
}
