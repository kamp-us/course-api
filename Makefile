gen:
	# Auto-generate code
	protoc --twirp_out=. --go_out=. rpc/course-api/service.proto