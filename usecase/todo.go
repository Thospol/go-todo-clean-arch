package usecase

import (
	"github.com/Thospol/go-todo-clean-arch/domain"
)

type todoUseCase struct {
	todoRepo domain.ToDoRepository
}

// NewToDoUseCase new instance of todo use case
func NewToDoUseCase(repo domain.ToDoRepository) domain.ToDoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}

// GetAllTodo get all todo
func (t *todoUseCase) GetAllTodo() (todos []domain.Todo, err error) {
	var entities []domain.Todo
	err = t.todoRepo.GetAllTodo(&entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// CreateTodo create todo
func (t *todoUseCase) CreateTodo(input *domain.Todo) (err error) {
	err = t.todoRepo.CreateTodo(input)
	if err != nil {
		return err
	}

	return nil
}

// GetTodo get todo
func (t *todoUseCase) GetTodo(input *domain.Todo, id string) (err error) {
	err = t.todoRepo.GetTodo(input, id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodo update todo
func (t *todoUseCase) UpdateTodo(input *domain.Todo, id string) (err error) {
	var checkingTodo domain.Todo
	err = t.todoRepo.GetTodo(&checkingTodo, id)
	if err != nil {
		return err
	}

	err = t.todoRepo.UpdateTodo(input, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo delete todo
func (t *todoUseCase) DeleteTodo(input *domain.Todo, id string) (err error) {
	err = t.todoRepo.DeleteTodo(input, id)
	if err != nil {
		return err
	}
	return nil
}
