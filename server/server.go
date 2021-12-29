package server

import (
	"context"
	"fmt"

	api "github.com/kamp-us/course-api/rpc/course-api"
)

type CourseAPIServer struct{}

func (s *CourseAPIServer) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (*api.Course, error) {
	fmt.Println(ctx, req, "<<<<<<<<<")
	return nil, nil
}
