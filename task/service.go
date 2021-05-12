package task

type Service interface {
	Index() ([]Task, error)
	Store(input InputTask) (Task, error)
	Show(id InputTaskDetail) (Task, error)
	Update(inputDetail InputTaskDetail, input InputTask) (Task, error)
	Destroy(id InputTaskDetail) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]Task, error) {
	tasks, err := s.repository.SelectAll()
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *service) Store(input InputTask) (Task, error) {
	task := Task{}
	task.Name = input.Name
	task.Description = input.Description

	newTask, err := s.repository.Insert(task)
	if err != nil {
		return newTask, err
	}
	return newTask, nil
}

func (s *service) Show(id InputTaskDetail) (Task, error) {
	task, err := s.repository.SelectById(id.ID)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *service) Update(inputDetail InputTaskDetail, input InputTask) (Task, error) {
	task, err := s.repository.SelectById(inputDetail.ID)
	if err != nil {
		return task, err
	}
	task.Name = input.Name
	task.Description = input.Description

	updatedTask, err := s.repository.Update(task)
	if err != nil {
		return updatedTask, err
	}
	return updatedTask, nil
}

func (s *service) Destroy(taskDetail InputTaskDetail) (bool, error) {
	_, err := s.repository.Destroy(taskDetail)
	if err != nil {
		return false, err
	}
	return true, nil
}
