package todo

import "gorm.io/gorm"

type Repository interface {
	Insert(todo Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}
