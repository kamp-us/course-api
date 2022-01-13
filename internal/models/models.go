package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Slug      string    `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string
	Description string

	UserID string `gorm:"index"`

	Categories []*CourseCategory
	Lessons    []*Lesson
}

type CourseCategory struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`

	CourseID   uuid.UUID
	CategoryID string
}

type Video struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Slug      string    `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID string `gorm:"index"`

	Name string
	URI  string
}

type Lesson struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Slug      string    `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string
	Description string

	UserID string `gorm:"index"`

	CourseID uuid.UUID
	Course   Course

	VideoID uuid.UUID
	Video   Video

	Categories []*LessonCategory
}

type LessonCategory struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	LessonID   uuid.UUID
	CategoryID string
}

func AutoMigrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&Course{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Video{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Lesson{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&CourseCategory{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&LessonCategory{}); err != nil {
		return err
	}
	return nil
}
