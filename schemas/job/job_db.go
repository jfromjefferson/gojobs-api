package job

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewJobDB(db *gorm.DB) *DB {
	return &DB{
		DB: db,
	}
}

func (db *DB) First(uuid uuid.UUID) (*Job, error) {
	var job Job
	err := db.DB.Where("uuid = ?", uuid).First(&job).Error
	if err != nil {
		return nil, err
	}

	return &job, err
}

func (db *DB) FindAll(title, sort string, page, limit int) ([]Job, error) {
	var jobs []Job
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = db.DB.Where("title LIKE ?", "%"+title+"%").Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&jobs).Error
	} else {
		err = db.DB.Where("title LIKE ?", "%"+title+"%").Order("created_at " + sort).Find(&jobs).Error
	}

	return jobs, err
}

func (db *DB) Create(job *Job) error {
	return db.DB.Create(job).Error
}

func (db *DB) Update(job *Job) error {
	_, err := db.First(job.Uuid)
	if err != nil {
		return err
	}

	return db.DB.Save(job).Error
}

func (db *DB) Delete(job *Job) error {
	_, err := db.First(job.Uuid)
	if err != nil {
		return err
	}

	return db.DB.Delete(job).Error
}
