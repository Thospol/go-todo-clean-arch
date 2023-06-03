package usecases

import (
	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// UseCase don't give a shit about the world!!
type todoUseCase struct {
	todoRepo domains.ToDoRepository
}

// NewToDoUseCase ...
func NewToDoUseCase(repo domains.ToDoRepository) domains.ToDoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}

func (t *todoUseCase) GetAllTodo() (todos []models.Todo, err error) {
	var todo []models.Todo
	handleErr := t.todoRepo.GetAllTodo(&todo)
	return todo, handleErr
}

func (t *todoUseCase) CreateTodo(input *models.Todo) (err error) {
	handleErr := t.todoRepo.CreateTodo(input)
	return handleErr
}

func (t *todoUseCase) GetTodo(input *models.Todo, id string) (err error) {
	handleErr := t.todoRepo.GetTodo(input, id)
	return handleErr
}

func (t *todoUseCase) UpdateTodo(input *models.Todo, id string) (err error) {
	// check avaliable
	var checkingTodo models.Todo
	errResp := t.todoRepo.GetTodo(&checkingTodo, id)
	if errResp != nil {
		return errResp
	}
	// update
	handleErr := t.todoRepo.UpdateTodo(input, id)
	return handleErr
}

func (t *todoUseCase) DeleteTodo(input *models.Todo, id string) (err error) {
	handleErr := t.todoRepo.DeleteTodo(input, id)
	return handleErr
}
