package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) GetCoursesByCategoryID(ctx context.Context, req *api.GetCoursesByCategoryIDRequest) (*api.GetCoursesByCategoryIDResponse, error) {
	if err := validateGetCoursesByCategoryIDRequest(req); err != nil {
		return nil, err
	}

	courses, err := s.backend.GetCoursesByCategoryID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	var batch []*api.Course
	for _, model := range courses {
		course := &api.Course{
			ID:          model.ID.String(),
			UserId:      model.UserID,
			Name:        model.Name,
			Description: model.Description,
			Slug:        model.Slug,
			CategoryIds: model.GetCategoryIDs(),
		}
		batch = append(batch, course)
	}

	return &api.GetCoursesByCategoryIDResponse{Courses: batch}, nil

}

func validateGetCoursesByCategoryIDRequest(req *api.GetCoursesByCategoryIDRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
