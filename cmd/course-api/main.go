package main

import (
	"net/http"

	course_api "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server"
)

func main() {
	server := &server.CourseAPIServer{}
	twirpHandler := course_api.NewCourseAPIServer(server)

	http.ListenAndServe(":80", twirpHandler)
}
