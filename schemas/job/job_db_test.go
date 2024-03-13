package job

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jfromjefferson/gojobs-api/internal/dto"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestDB_First(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(Job{})
	assert.Nil(t, err)

	jobDTO := dto.JobDTO{
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

	jobDB := NewJobDB(db)
	assert.NotNil(t, jobDB)

	err = jobDB.Create(job)
	assert.Nil(t, err)

	filteredJob, err := jobDB.First(job.Uuid)
	assert.NotNil(t, filteredJob)
	assert.Nil(t, err)
	assert.Equal(t, filteredJob.Uuid, job.Uuid)

	filteredJob, err = jobDB.First(uuid.New())
	assert.Nil(t, filteredJob)
	assert.NotNil(t, err)
}

func TestDB_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(Job{})
	assert.Nil(t, err)

	jobDB := NewJobDB(db)
	assert.NotNil(t, jobDB)

	for i := 1; i <= 20; i++ {
		jobDTO := dto.JobDTO{
			Title:       fmt.Sprintf("Title %d", i),
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

		err = jobDB.Create(job)
		assert.Nil(t, err)
	}

	jobs, err := jobDB.FindAll("asc", 1, 10)
	assert.Nil(t, err)
	assert.NotNil(t, jobs)
	assert.Len(t, jobs, 10)

	jobs, err = jobDB.FindAll("asc", 150, 10)
	assert.NotNil(t, jobs)
	assert.Nil(t, err)
	assert.Len(t, jobs, 0)

}

func TestDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(Job{})
	assert.Nil(t, err)

	jobDTO := dto.JobDTO{
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

	jobDB := NewJobDB(db)
	assert.NotNil(t, jobDB)

	err = jobDB.Create(job)
	assert.Nil(t, err)

	jobDTO.Salary = -1
	job, err = NewJob(jobDTO)
	assert.Nil(t, job)
	assert.NotNil(t, err)
}

func TestDB_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(Job{})
	assert.Nil(t, err)

	jobDB := NewJobDB(db)
	assert.NotNil(t, jobDB)

	jobDTO := dto.JobDTO{
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

	err = jobDB.Create(job)
	assert.Nil(t, err)

	filteredJob, err := jobDB.First(job.Uuid)
	assert.Nil(t, err)
	assert.NotNil(t, filteredJob)

	filteredJob.Title = "Test"
	err = jobDB.Update(filteredJob)
	assert.Nil(t, err)
	filteredJob, err = jobDB.First(job.Uuid)
	assert.Nil(t, err)
	assert.NotNil(t, filteredJob)
	assert.Equal(t, "Test", filteredJob.Title)
}

func TestDB_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(Job{})
	assert.Nil(t, err)

	jobDB := NewJobDB(db)
	assert.NotNil(t, jobDB)

	jobDTO := dto.JobDTO{
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

	err = jobDB.Create(job)
	assert.Nil(t, err)

	filteredJob, err := jobDB.First(job.Uuid)
	assert.Nil(t, err)
	assert.NotNil(t, filteredJob)

	err = jobDB.Delete(filteredJob)
	assert.Nil(t, err)
	jobs, err := jobDB.FindAll("asc", 1, 10)
	assert.NotNil(t, jobs)
	assert.Nil(t, err)
	assert.Len(t, jobs, 0)
}
