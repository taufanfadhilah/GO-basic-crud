package todo

type Service interface {
	Store(input InputTodo) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input InputTodo) (Todo, error) {
	todo := Todo{}
	todo.Name = input.Name
	todo.Description = input.Description

	newTodo, err := s.repository.Insert(todo)
	if err != nil {
		return newTodo, err
	}
	return newTodo, nil
}
