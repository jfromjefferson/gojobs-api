package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jfromjefferson/gojobs-api/internal/dto"
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"net/http"
)

// UpdateJobHandler Update job godoc
// @Summary Update job
// @Description Update job...
// @Tags jobs
// @Accept json
// @Produce json
// @Param uuid query string true "uuid" Format(uuid)
// @Param request body dto.JobDTO true "job request"
// @Success 200
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /jobs/job [put]
func UpdateJobHandler(context *gin.Context) {
	uuidParam := context.Query("uuid")
	if uuidParam == "" {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "uuid param is required",
		})
		return
	}

	jobDB := job.NewJobDB(db)

	uuidParse, err := uuid.Parse(uuidParam)
	if err != nil {
		errMessage := fmt.Sprintf("Parse uuid error: %v", err)
		logger.ErrF("Parse uuid error: ", errMessage)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": errMessage,
		})
		return
	}

	filteredJob, err := jobDB.First(uuidParse)
	if err != nil {
		logger.ErrF("Error", err)
		context.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	var jobDTO dto.JobDTO
	err = json.NewDecoder(context.Request.Body).Decode(&jobDTO)
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		logger.Err(errMessage)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	filteredJob.Title = jobDTO.Title
	filteredJob.Description = jobDTO.Description
	filteredJob.Company = jobDTO.Company
	filteredJob.Location = jobDTO.Location
	filteredJob.Salary = jobDTO.Salary
	filteredJob.IsActive = jobDTO.IsActive
	filteredJob.IsRemote = jobDTO.IsRemote
	filteredJob.Link = jobDTO.Link

	err = filteredJob.Validate()
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		logger.Err(errMessage)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	err = jobDB.Update(filteredJob)
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		logger.Err(errMessage)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Job updated successfully",
	})
}
