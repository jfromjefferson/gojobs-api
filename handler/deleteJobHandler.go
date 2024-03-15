package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"net/http"
)

// DeleteJobHandler Delete job godoc
// @Summary Delete job
// @Description Delete job...
// @Tags jobs
// @Accept json
// @Produce json
// @Param uuid query string true "uuid" Format(uuid)
// @Success 200
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /jobs/job [delete]
func DeleteJobHandler(context *gin.Context) {
	uuidParam := context.Query("uuid")
	if uuidParam == "" {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "uuid param is required",
		})
		return
	}

	uuidParse, err := uuid.Parse(uuidParam)
	if err != nil {
		errMessage := fmt.Sprintf("Parse uuid error: %v", err)
		logger.ErrF("Parse uuid error: ", errMessage)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": errMessage,
		})
		return
	}

	jobDB := job.NewJobDB(db)
	filteredJob, err := jobDB.First(uuidParse)
	if err != nil {
		logger.ErrF("Error", err)
		context.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	err = jobDB.Delete(filteredJob)
	if err != nil {
		errMessage := fmt.Sprintf("Error : %v", err)
		logger.Err(errMessage)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Job deleted successfully",
	})
}
