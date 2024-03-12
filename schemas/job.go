package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Salary      int64  `json:"salary"`
	IsActive    bool   `json:"isActive"`
	IsRemote    bool   `json:"isRemote"`
	Link        string `json:"link"`
	Uuid        uuid.UUID
}
