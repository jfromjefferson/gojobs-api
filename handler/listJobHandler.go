package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
