package server

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"gorm.io/gorm"
)

type CourseAPIServer struct {
	Db *gorm.DB
}

func (s *CourseAPIServer) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (*api.Course, error) {
	course := models.Course{
		Slug:        slug.MakeLang(req.Name, "tr"),
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserID,
	}

	result := s.Db.Create(&course)
	if result.Error != nil {
		return nil, result.Error
	}

	return &api.Course{
		ID:          course.ID.String(),
		UserID:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
		Slug:        course.Slug,
	}, nil

}
