package router

import "github.com/gin-gonic/gin"

func Init(addr string) {
	router := gin.Default()

	router.GET("/api/v1/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hi",
		})
	})

	err := router.Run(addr)
	if err != nil {
		panic(err)
	}
}
