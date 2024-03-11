package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func GetOneJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func CreateJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func UpdateJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

func DeleteJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
