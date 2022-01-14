package server

import (
	"github.com/kamp-us/course-api/internal/backend"
)

type CourseAPIServer struct {
	backend backend.Backender
}

func NewCourseAPIServer(backend backend.Backender) *CourseAPIServer {
	return &CourseAPIServer{backend: backend}
}
