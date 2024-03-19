package job

import (
	"github.com/jfromjefferson/gojobs-api/internal/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJob(t *testing.T) {
	jobDTO := dto.JobDTOInput{
		Title:       "Full-stack Developer",
		Description: "Job description",
		Company:     "Meta Platforms",
		Location:    "Palo Alto - California",
		Salary:      1500000,
		IsActive:    true,
		IsRemote:    true,
		Link:        "https://meta.com",
	}
	job, err := NewJob(jobDTO)

	assert.Nil(t, err)
	assert.NotNil(t, job)
	assert.NotEmpty(t, job.Title)
	assert.NotEmpty(t, job.Uuid)
	assert.True(t, job.Validate() == nil)

	jobDTO2 := dto.JobDTOInput{
		Description: "Job description",
		Company:     "Meta Platforms",
		Location:    "Palo Alto - California",
		Salary:      1500000,
		Link:        "https://meta.com",
	}
	job, err = NewJob(jobDTO2)
	assert.Nil(t, job)
	assert.NotNil(t, err)
	assert.Equal(t, ErrTitleIsRequired, err)
}
