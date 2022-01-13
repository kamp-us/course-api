package server

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

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
