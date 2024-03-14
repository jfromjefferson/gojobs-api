package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"net/http"
)

func GetOneJobHandler(context *gin.Context) {
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

	context.JSON(http.StatusOK, gin.H{"data": filteredJob})
}
