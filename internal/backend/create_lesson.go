package backend

import (
	"context"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) CreateLesson(ctx context.Context, userID string, name string, description string, courseID uuid.UUID, categoryIDs []string) (*models.Lesson, error) {
	lesson := models.Lesson{
		Slug:        slug.MakeLang(name, "tr"),
		Name:        name,
		Description: description,
		UserID:      userID,
		CourseID:    courseID,
	}

	result := b.DB.Create(&lesson)
	if result.Error != nil {
		return nil, result.Error
	}

	var categories []*models.LessonCategory
	for _, categoryId := range categoryIDs {
		lessonCategory := models.LessonCategory{LessonID: lesson.ID, CategoryID: categoryId}
		categories = append(categories, &lessonCategory)
	}

	err := b.DB.Model(&lesson).Association("Categories").Append(categories)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}
