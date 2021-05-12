package task

import "gorm.io/gorm"

type Repository interface {
	Insert(task Task) (Task, error)
	SelectAll() ([]Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *repository) SelectAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
