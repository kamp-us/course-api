package server

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func (s *CourseAPIServer) UpdateCourse(ctx context.Context, req *api.UpdateCourseRequest) (*api.Course, error) {
	course := models.Course{}
	result := s.Db.First(&course, "id = ?", req.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	updates := models.Course{}

	if name := convertToStringPtr(req.Name); name != nil {
		updates.Name = *name
		updates.Slug = slug.MakeLang(*name, "tr")
	}
	if description := convertToStringPtr(req.Description); description != nil {
		updates.Description = *description
	}

	result = s.Db.Model(&course).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return &api.Course{
		ID:          course.ID.String(),
		UserId:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
		Slug:        course.Slug,
	}, nil
}
