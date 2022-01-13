package server

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
)

func (s *CourseAPIServer) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (*api.Course, error) {
	course := models.Course{
		Slug:        slug.MakeLang(req.Name, "tr"),
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserId,
	}

	result := s.Db.Create(&course)
	if result.Error != nil {
		return nil, result.Error
	}

	//categoryIds := make([]string, 0)
	//for _, categoryId := range req.CategoryIds {
	//	courseCategory := models.CourseCategory{CourseID: course.ID, CategoryID: categoryId}
	//	query := s.Db.Create(&courseCategory)
	//	if query.Error != nil {
	//		return nil, query.Error
	//	}
	//	categoryIds = append(categoryIds, categoryId)
	//}

	return &api.Course{
		ID:          course.ID.String(),
		UserId:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
		Slug:        course.Slug,
		//CategoryIds: categoryIds,
	}, nil
}
