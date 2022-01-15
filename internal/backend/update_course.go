package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
)

func (b *PostgreSQLBackend) UpdateCourse(ctx context.Context, id string, name *string, description *string) error {
	course := models.Course{}
	result := b.DB.First(&course, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	updates := models.Course{Name: *name, Slug: slug.MakeLang(*name, "tr"), Description: *description}
	result = b.DB.Model(&course).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
