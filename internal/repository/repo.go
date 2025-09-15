package repository

import (
	"database/sql"
)

type Task struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	Priority  string `json:"priority"`
	DueDate   string `json:"due_date"`
}

type TaskRepository interface {
	AddTask(task Task) (int, error)
	GetTasks() ([]Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
}

type SQLiteTaskRepository struct {
	DB *sql.DB
}

func (r *SQLiteTaskRepository) AddTask(task Task) (int, error) {
	statement, err := r.DB.Prepare("INSERT INTO tasks (text, completed, priority, due_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := statement.Exec(task.Text, task.Completed, task.Priority, task.DueDate)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *SQLiteTaskRepository) GetTasks() ([]Task, error) {
	rows, err := r.DB.Query("SELECT id, text, completed, priority, due_date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.Priority, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *SQLiteTaskRepository) UpdateTask(task Task) error {
	statement, err := r.DB.Prepare("UPDATE tasks SET text = ?, completed = ?, priority = ?, due_date = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(task.Text, task.Completed, task.Priority, task.DueDate, task.ID)
	return err
}

func (r *SQLiteTaskRepository) DeleteTask(id int) error {
	statement, err := r.DB.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}
