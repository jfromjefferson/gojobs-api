package router

import "github.com/gin-gonic/gin"

func Init(addr string) {
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(addr)
	if err != nil {
		panic(err)
	}
}
