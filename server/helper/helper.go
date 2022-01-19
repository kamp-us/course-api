package helper

import (
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func ConvertToLessonModel(l *models.Lesson) *api.Lesson {
	return &api.Lesson{
		ID:          l.ID.String(),
		UserId:      l.UserID,
		CourseId:    l.CourseID.String(),
		Name:        l.Name,
		Description: l.Description,
		Slug:        l.Slug,
		CategoryIds: l.GetCategoryIDs(),
	}
}

func ConvertToCourseModel(c *models.Course) *api.Course {
	return &api.Course{
		ID:          c.ID.String(),
		UserId:      c.UserID,
		Name:        c.Name,
		Description: c.Description,
		Slug:        c.Slug,
		CategoryIds: c.GetCategoryIDs(),
	}
}
