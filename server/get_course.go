package server

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func (s *CourseAPIServer) GetCourse(ctx context.Context, req *api.GetCourseRequest) (*api.Course, error) {
	course := models.Course{}
	result := s.Db.First(&course, "id = ?", req.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	var categories []models.CourseCategory
	query := s.Db.Model(&course).Association("Categories").Find(&categories)
	if query != nil {
		return nil, query
	}

	categoryIds := make([]string, 0)
	for _, category := range categories {
		categoryIds = append(categoryIds, category.CategoryID)
	}

	return &api.Course{
		ID:          course.ID.String(),
		UserId:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
		Slug:        course.Slug,
		CategoryIds: categoryIds,
	}, nil

}
