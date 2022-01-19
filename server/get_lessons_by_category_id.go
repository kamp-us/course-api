package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) GetLessonsByCategoryID(ctx context.Context, req *api.GetLessonsByCategoryIDRequest) (*api.GetLessonsByCategoryIDResponse, error) {
	if err := validateGetLessonsByCategoryIDRequest(req); err != nil {
		return nil, err
	}

	lessons, err := s.backend.GetLessonsByCategoryID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	var batch []*api.Lesson
	for _, model := range lessons {
		course := &api.Lesson{
			ID:          model.ID.String(),
			UserId:      model.UserID,
			CourseId:    model.CourseID.String(),
			Name:        model.Name,
			Description: model.Description,
			Slug:        model.Slug,
			CategoryIds: model.GetCategoryIDs(),
		}
		batch = append(batch, course)
	}

	return &api.GetLessonsByCategoryIDResponse{Lessons: batch}, nil

}

func validateGetLessonsByCategoryIDRequest(req *api.GetLessonsByCategoryIDRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
