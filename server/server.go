package server

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
		UserID:      req.UserId,
	}

	result := s.Db.Create(&course)
	if result.Error != nil {
		return nil, result.Error
	}

	categoryIds := make([]string, 0)

	for _, categoryId := range req.CategoryIds {
		courseCategory := models.CourseCategory{CourseID: course.ID, CategoryID: categoryId}
		query := s.Db.Create(&courseCategory)
		if query.Error != nil {
			return nil, query.Error
		}
		categoryIds = append(categoryIds, categoryId)
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

func (s *CourseAPIServer) GetCourse(ctx context.Context, req *api.GetCourseRequest) (*api.Course, error) {
	course := models.Course{}
	result := s.Db.First(&course, "id = ?", req.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	//categories := models.CourseCategory{}
	//result = s.Db.Where("course_id = ?", req.ID).Find(&categories)

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

func (s *CourseAPIServer) DeleteCourse(ctx context.Context, req *api.DeleteCourseRequest) (*emptypb.Empty, error) {
	course := models.Course{}
	result := s.Db.First(&course, "id = ?", req.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	result = s.Db.Delete(&course)
	if result.Error != nil {
		return nil, result.Error
	}

	return &emptypb.Empty{}, nil
}

func convertToStringPtr(value *wrapperspb.StringValue) *string {
	if value == nil {
		return nil
	}
	val := value.GetValue()
	return &val
}
