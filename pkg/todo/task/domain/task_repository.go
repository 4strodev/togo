package domain

type TaskRepository interface {
	GetTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	SaveTask(task Task) error
	UpdateTask(task Task) (Task, error)
	DeleteTask(id string) error
}
