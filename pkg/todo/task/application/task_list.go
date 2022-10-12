package application

import "github.com/4strodev/todo_app/pkg/todo/task/domain"

type TaskList struct {
	repository domain.TaskRepository
}

func (self *TaskList) SetRepository(repository domain.TaskRepository) {
	self.repository = repository
}

func (self *TaskList) Run() ([]domain.Task, error) {
	tasks := []domain.Task{}

	tasks, err := self.repository.GetTasks()

	return tasks, err
}
