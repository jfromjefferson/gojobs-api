package dto

import (
	"github.com/google/uuid"
	"time"
)

type JobDTOInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Salary      int64  `json:"salary"`
	IsActive    bool   `json:"isActive"`
	IsRemote    bool   `json:"isRemote"`
	Link        string `json:"link"`
}

type JobDTOOutput struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Salary      int64     `json:"salary"`
	IsActive    bool      `json:"isActive"`
	IsRemote    bool      `json:"isRemote"`
	Link        string    `json:"link"`
	Uuid        uuid.UUID `json:"uuid"`
}
