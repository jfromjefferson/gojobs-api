package main

import (
	configs "github.com/jfromjefferson/gojobs-api/config"
	_ "github.com/jfromjefferson/gojobs-api/docs"
	"github.com/jfromjefferson/gojobs-api/handler"
	"github.com/jfromjefferson/gojobs-api/router"
)

func main() {
	logger := configs.GetLogger("main")
	// Init database
	_, err := configs.LoadConfig(".")
	if err != nil {
		logger.ErrF("Config init error: %v", err)
		return
	}

	// Init database
	db, err := configs.InitDatabase()
	if err != nil {
		logger.ErrF("Database init error: %v", err)
	}

	handler.InitHandler(db)

	// Init routes
	router.Init(":8000")
}
