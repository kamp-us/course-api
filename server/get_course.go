package server

import (
	"context"
	api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/twitchtv/twirp"
)

func (s *CourseAPIServer) GetCourse(ctx context.Context, req *api.GetCourseRequest) (*api.Course, error) {
	if err := validateGetCourseRequest(req); err != nil {
		return nil, err
	}

	course, err := s.backend.GetCourse(ctx, req.ID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &api.Course{
		ID:          course.ID.String(),
		UserId:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
		Slug:        course.Slug,
		CategoryIds: course.GetCategoryIDs(),
	}, nil

}

func validateGetCourseRequest(req *api.GetCourseRequest) error {
	if req.ID == "" {
		return twirp.RequiredArgumentError("ID")
	}
	return nil
}
