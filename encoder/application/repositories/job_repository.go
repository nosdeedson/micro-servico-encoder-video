package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func (r JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	err := r.Db.Create(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (r JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
	r.Db.Preload("Video").First(&job, "id =?", id)
	if job.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}
	return &job, nil
}

func (r JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := r.Db.Save(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}
