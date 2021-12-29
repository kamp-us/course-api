package main

import (
	"net/http"

	courseapi "github.com/kamp-us/course-api/rpc/course-api"
	"github.com/kamp-us/course-api/server"
)

func main() {
	server := &server.CourseAPIServer{}
	twirpHandler := courseapi.NewCourseAPIServer(server)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	mux.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("OK"))
	}))

	http.ListenAndServe(":80", mux)
}
