package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOneJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
