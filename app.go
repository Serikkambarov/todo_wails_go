package main

import (
	"context"
	"database/sql"
	"myproject/internal/repository"
)

// App struct
type App struct {
	ctx      context.Context
	taskRepo repository.TaskRepository
}

// NewApp creates a new App application struct
func NewApp(db *sql.DB) *App {
	return &App{taskRepo: &repository.SQLiteTaskRepository{DB: db}}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetTasks() ([]repository.Task, error) {
	return a.taskRepo.GetTasks()
}

func (a *App) AddTask(task repository.Task) (int, error) {
	return a.taskRepo.AddTask(task)
}

func (a *App) UpdateTask(task repository.Task) error {
	return a.taskRepo.UpdateTask(task)
}

func (a *App) DeleteTask(id int) error {
	return a.taskRepo.DeleteTask(id)
}
