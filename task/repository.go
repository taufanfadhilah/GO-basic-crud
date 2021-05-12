package task

import "gorm.io/gorm"

type Repository interface {
	Insert(task Task) (Task, error)
	SelectAll() ([]Task, error)
	SelectById(id int) (Task, error)
	Update(task Task) (Task, error)
	Destroy(taskDetail InputTaskDetail) (bool, error)
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

func (r *repository) SelectById(id int) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *repository) Update(task Task) (Task, error) {
	err := r.db.Save(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *repository) Destroy(taskDetail InputTaskDetail) (bool, error) {
	task := Task{
		ID: taskDetail.ID,
	}
	err := r.db.Delete(&task).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
