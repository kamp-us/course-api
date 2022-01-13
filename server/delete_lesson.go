package server

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CourseAPIServer) DeleteLesson(ctx context.Context, req *api.DeleteLessonRequest) (*emptypb.Empty, error) {
	lesson := models.Lesson{}
	result := s.Db.First(&lesson, "id = ?", req.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	result = s.Db.Delete(&lesson)

	if result.Error != nil {
		return nil, result.Error
	}

	return &emptypb.Empty{}, nil
}
