package models

import (
	"time"

	"github.com/google/uuid"
)

type Bucket struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title       string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Versioning  bool        `gorm:"not null" json:"content,omitempty"`
	Locking     bool        `gorm:"not null" json:"image,omitempty"`
	Quota       bool        `gorm:"not null" json:"user,omitempty"`
	User        string      `json:"user,omitempty"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateBucketRequest struct {
	Title       string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Versioning  bool        `gorm:"not null" json:"content,omitempty"`
	Locking     bool        `gorm:"not null" json:"image,omitempty"`
	Quota       bool        `gorm:"not null" json:"user,omitempty"`
	User        string      `json:"user,omitempty"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}

type UpdateBucket struct {
	Title       string      `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Versioning  bool        `gorm:"not null" json:"content,omitempty"`
	Locking     bool        `gorm:"not null" json:"image,omitempty"`
	Quota       bool        `gorm:"not null" json:"user,omitempty"`
	User        string      `json:"user,omitempty"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at,omitempty"`
}
