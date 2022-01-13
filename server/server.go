package server

import (
	"gorm.io/gorm"
)

type CourseAPIServer struct {
	Db *gorm.DB
}
