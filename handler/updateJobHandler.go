package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
