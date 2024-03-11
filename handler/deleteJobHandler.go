package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
