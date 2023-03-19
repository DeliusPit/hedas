package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title       string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Bucket      uuid.UUID
	User        uuid.UUID    `json:"user,omitempty"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateFileRequest struct {
	Title       string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	User        string      `json:"user,omitempty"`
	Bucket      string
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}

type UpdateFile struct {
	Title       string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	User        string      `json:"user,omitempty"`
	Bucket      string
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}
