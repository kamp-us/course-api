package backend

import (
	"context"
	"github.com/kamp-us/course-api/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	GetCourse(ctx context.Context, id string) (*models.Course, error)
	CreateCourse(ctx context.Context, userID string, name string, description string) (*models.Course, error)
	UpdateCourse(ctx context.Context, id string, name *string, description *string) error
	DeleteCourse(ctx context.Context, id string) error
	GetLesson(ctx context.Context, id string) (*models.Lesson, error)
	CreateLesson(ctx context.Context, userID string, name string, description string, courseID string) (*models.Lesson, error)
	UpdateLesson(ctx context.Context, id string, name *string, description *string) error
	DeleteLesson(ctx context.Context, id string) error
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}
