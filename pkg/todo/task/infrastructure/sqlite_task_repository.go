package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/4strodev/todo_app/pkg/todo/task/domain"
)

type SqliteTaskRepository struct {
	pool *sql.DB
}

// set the pool that repository will use
func (self *SqliteTaskRepository) SetPool(pool *sql.DB) {
	self.pool = pool
}

// Get all tasks
func (self *SqliteTaskRepository) GetTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	rows, err := self.pool.Query("SELECT id, user_id, name, description FROM tasks")
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return []domain.Task{}, errors.New("Error getting tasks")
	}

	var task = domain.Task{}
	for rows.Next() {
		err = rows.Scan(&task.Id, &task.UserId, &task.Name, &task.Description)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Error populating tasks")
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Get task based on their id
func (self *SqliteTaskRepository) GetTask(id string) (domain.Task, error) {
	var task domain.Task

	rows, err := self.pool.Query("SELECT id, user_id, name, description FROM tasks WHERE id = ?")
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return task, fmt.Errorf("Error getting task identified with %s", id)
	}

	for rows.Next() {
		err := rows.Scan(&task.Id, &task.UserId, &task.Name, &task.Description)
		if err != nil {
			log.Println(err)
			return domain.Task{}, fmt.Errorf("Error populating task identified with %s", id)
		}
	}

	return task, nil
}

// Save new task
func (self *SqliteTaskRepository) SaveTask(task domain.Task) error {
	insertQuery := "INSERT INTO tasks (id, user_id, name, description) VALUES (?, ?, ?, ?)"
	rows, err := self.pool.Query(insertQuery, task.Id, task.UserId, task.Name, task.Description)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return errors.New("Error saving new task")
	}

	return nil
}

// Update task information
func (self *SqliteTaskRepository) UpdateTask(task domain.Task) (domain.Task, error) {
	updateQuery := "UPDATE tasks set user_id = ?, name = ?, description = ? WHERE id = ?"

	rows, err := self.pool.Query(updateQuery, task.UserId, task.Name, task.Description, task.Id)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return domain.Task{}, errors.New("Error updating task")
	}

	return task, nil
}

// Delete task based on their id
func (self *SqliteTaskRepository) DeleteTask(id string) error {
	deleteQuery := "DELETE FROM tasks WHERE id = ?"

	rows, err := self.pool.Query(deleteQuery, id)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return errors.New("Error deleting new task")
	}

	return nil
}
