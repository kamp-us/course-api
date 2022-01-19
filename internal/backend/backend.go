package backend

import (
	"context"
	"github.com/google/uuid"
	"github.com/kamp-us/course-api/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	GetCourse(ctx context.Context, id string) (*models.Course, error)
	CreateCourse(ctx context.Context, userID string, name string, description string, categoryIDs []string) (*models.Course, error)
	UpdateCourse(ctx context.Context, id string, name *string, description *string) error
	DeleteCourse(ctx context.Context, id string) error
	GetBatchCourses(ctx context.Context, ids []string) ([]*models.Course, error)
	GetCoursesByCategoryID(ctx context.Context, id string) ([]*models.Course, error)
	GetLesson(ctx context.Context, id string) (*models.Lesson, error)
	CreateLesson(ctx context.Context, userID string, name string, description string, courseID uuid.UUID, categoryIDs []string) (*models.Lesson, error)
	UpdateLesson(ctx context.Context, id string, name *string, description *string) error
	DeleteLesson(ctx context.Context, id string) error
	GetBatchLessons(ctx context.Context, ids []string) ([]*models.Lesson, error)
	GetLessonsByCategoryID(ctx context.Context, id string) ([]*models.Lesson, error)
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}
