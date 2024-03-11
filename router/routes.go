package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/jobs", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /jobs",
			})
		})

		v1.GET("/jobs/job", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /jobs/job",
			})
		})

		v1.POST("/jobs", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST /jobs/job",
			})
		})

		v1.PUT("/jobs/job", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "PUT /jobs/job",
			})
		})

		v1.DELETE("/jobs/job", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "DELETE /jobs/job",
			})
		})
	}
}
