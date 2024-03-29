package job

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jfromjefferson/gojobs-api/internal/dto"
	"gorm.io/gorm"
)

var (
	ErrTitleIsRequired       = errors.New("title is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrCompanyIsRequired     = errors.New("company is required")
	ErrLocationIsRequired    = errors.New("location is required")
	ErrSalaryIsRequired      = errors.New("salary is required")
	ErrLinkIsRequired        = errors.New("link is required")
)

type Job struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
	IsActive    bool      `json:"isActive"`
	IsRemote    bool      `json:"isRemote"`
	Uuid        uuid.UUID `json:"uuid"`
}

func NewJob(dto dto.JobDTOInput) (*Job, error) {
	job := Job{
		Title:       dto.Title,
		Description: dto.Description,
		Company:     dto.Company,
		Location:    dto.Location,
		Salary:      dto.Salary,
		IsActive:    dto.IsActive,
		IsRemote:    dto.IsRemote,
		Link:        dto.Link,
		Uuid:        uuid.New(),
	}

	err := job.Validate()
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (job *Job) Validate() error {
	if job.Title == "" {
		return ErrTitleIsRequired
	}

	if job.Description == "" {
		return ErrDescriptionIsRequired
	}

	if job.Company == "" {
		return ErrCompanyIsRequired
	}

	if job.Location == "" {
		return ErrLocationIsRequired
	}

	if job.Salary <= 0 {
		return ErrSalaryIsRequired
	}

	if job.Link == "" {
		return ErrLinkIsRequired
	}

	return nil
}
