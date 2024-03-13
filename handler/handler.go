package handler

import (
	configs "github.com/jfromjefferson/gojobs-api/config"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitHandler(database *gorm.DB) {
	logger = configs.GetLogger("handler")
	db = database
}
