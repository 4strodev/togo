package application

import (
	"fmt"

	"github.com/4strodev/todo_app/pkg/todo/task/domain"
	"github.com/google/uuid"
)

const TASK_SUFIX = "TSK"

type TaskCreate struct {
	repository domain.TaskRepository
}

func (self *TaskCreate) SetRepository(repository domain.TaskRepository) {
	self.repository = repository
}

func (self *TaskCreate) Run(userId string, name string, description string) error {
	task := domain.Task{}

	task.Id = newTaskId()
	task.UserId = userId
	task.Name = name
	task.Description = description

	return self.repository.SaveTask(task)
}

func newTaskId() string {
	return fmt.Sprintf("%s:%s", uuid.NewString(), TASK_SUFIX)
}
