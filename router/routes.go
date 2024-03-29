package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jfromjefferson/gojobs-api/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/jobs", handler.ListJobHandler)

		v1.GET("/jobs/job", handler.GetOneJobHandler)

		v1.POST("/jobs", handler.CreateJobHandler)

		v1.PUT("/jobs/job", handler.UpdateJobHandler)

		v1.DELETE("/jobs/job", handler.DeleteJobHandler)

		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
