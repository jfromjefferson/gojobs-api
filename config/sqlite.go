package configs

import (
	"github.com/jfromjefferson/gojobs-api/schemas/job"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(job.Job{})
	if err != nil {
		return nil, err
	}

	return db, err
}
