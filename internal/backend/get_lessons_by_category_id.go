package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) GetLessonsByCategoryID(ctx context.Context, id string) ([]*models.Lesson, error) {
	var lessons []*models.Lesson
	var lessonCategories []*models.LessonCategory

	query := b.DB.Preload("Lesson.Categories").Preload("Lesson").Find(&lessonCategories, "category_id = ?", id)
	if query.Error != nil {
		return nil, query.Error
	}

	for _, model := range lessonCategories {
		lessons = append(lessons, &model.Lesson)
	}

	return lessons, nil
}
