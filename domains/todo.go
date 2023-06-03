package domains

import (
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoUseCase ...
type ToDoUseCase interface {
	GetAllTodo() (t []models.Todo, err error)
	CreateTodo(t *models.Todo) (err error)
	GetTodo(t *models.Todo, id string) (err error)
	UpdateTodo(t *models.Todo, id string) (err error)
	DeleteTodo(t *models.Todo, id string) (err error)
}

// ToDoRepository ...
type ToDoRepository interface {
	GetAllTodo(t *[]models.Todo) (err error)
	CreateTodo(t *models.Todo) (err error)
	GetTodo(t *models.Todo, id string) (err error)
	UpdateTodo(t *models.Todo, id string) (err error)
	DeleteTodo(t *models.Todo, id string) (err error)
}
