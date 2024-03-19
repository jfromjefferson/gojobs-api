package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jfromjefferson/gojobs-api/internal/dto"
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

// CreateJobHandler Create job godoc
// @Summary Create job
// @Description Create job...
// @Tags jobs
// @Accept json
// @Produce json
// @Param request body dto.JobDTOInput true "job request"
// @Success 201
// @Failure 500 {object} Error
// @Router /jobs [post]
func CreateJobHandler(context *gin.Context) {
	var jobDTO dto.JobDTOInput
	err := json.NewDecoder(context.Request.Body).Decode(&jobDTO)
	if err != nil {
		errMessage := fmt.Sprintf("Error creating job: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		logger.Err(errMessage)
		return
	}

	jobDB := job.NewJobDB(db)

	newJob, err := job.NewJob(jobDTO)
	if err != nil {
		errMessage := fmt.Sprintf("Error creating job: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		logger.Err(errMessage)
		return
	}

	err = jobDB.Create(newJob)
	if err != nil {
		errMessage := fmt.Sprintf("Error saving job in database: %v", err)
		logger.Err(errMessage)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Job created successfully",
	})
}
