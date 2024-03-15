package dto

import (
	"github.com/google/uuid"
	"time"
)

type JobDTO struct {
	ID          uint      `json:"id" swaggerignore:"true"`
	CreatedAt   time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updatedAt" swaggerignore:"true"`
	DeletedAt   time.Time `json:"deletedAt,omitempty" swaggerignore:"true"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Salary      int64     `json:"salary"`
	IsActive    bool      `json:"isActive"`
	IsRemote    bool      `json:"isRemote"`
	Link        string    `json:"link"`
	Uuid        uuid.UUID `json:"uuid" swaggerignore:"true"`
}
