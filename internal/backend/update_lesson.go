package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) UpdateLesson(ctx context.Context, id string, name *string, description *string) error {
	lesson := models.Lesson{}
	if result := b.DB.First(&lesson, "id = ?", id); result.Error != nil {
		return result.Error
	}

	updates := models.Lesson{Name: *name, Slug: slug.MakeLang(*name, "tr"), Description: *description}

	if result := b.DB.Model(&lesson).Updates(updates); result.Error != nil {
		return result.Error
	}

	return nil
}
