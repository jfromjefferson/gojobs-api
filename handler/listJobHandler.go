package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"net/http"
	"strconv"
)

// ListJobHandler List job godoc
// @Summary List job
// @Description List job...
// @Tags jobs
// @Accept json
// @Produce json
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Param title query string false "title"
// @Success 200 {array} dto.JobDTOOutput
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /jobs [get]
func ListJobHandler(context *gin.Context) {
	page := context.Query("page")
	limit := context.Query("limit")
	sort := context.Query("sort")
	title := context.Query("title")

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		logger.Err(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		logger.Err(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	jobDB := job.NewJobDB(db)

	jobs, err := jobDB.FindAll(title, sort, pageInt, limitInt)
	if err != nil {
		errMessage := fmt.Sprintf("Error getting job list: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		logger.Err(errMessage)
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": jobs})
}
