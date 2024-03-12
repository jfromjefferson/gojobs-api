package main

import (
	configs "github.com/jfromjefferson/gojobs-api/config"
	"github.com/jfromjefferson/gojobs-api/router"
	"github.com/jfromjefferson/gojobs-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	logger := configs.GetLogger("main")
	// Init database
	_, err := configs.LoadConfig(".")
	if err != nil {
		logger.ErrF("Config init error: %v", err)
		return
	}

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(schemas.Job{})
	if err != nil {
		panic(err)
	}

	// Init routes
	router.Init(":8000")
}
