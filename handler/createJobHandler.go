package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jfromjefferson/gojobs-api/internal/dto"
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"net/http"
)

func CreateJobHandler(context *gin.Context) {
	var jobDTO dto.JobDTO
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
