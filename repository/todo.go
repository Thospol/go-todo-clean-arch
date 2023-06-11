package repository

import (
	"github.com/Thospol/go-todo-clean-arch/domain"
	_ "github.com/go-sql-driver/mysql" // use to connect db
	"github.com/jinzhu/gorm"
)

type todoRepository struct {
	conn *gorm.DB
}

// NewToDoRepository new instance of todo repository
func NewToDoRepository(conn *gorm.DB) domain.ToDoRepository {
	return &todoRepository{conn}
}

// GetAllTodo get all todo
func (t *todoRepository) GetAllTodo(todo *[]domain.Todo) (err error) {
	if err = t.conn.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateTodo create todo
func (t *todoRepository) CreateTodo(todo *domain.Todo) (err error) {
	err = t.conn.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

// GetTodo get todo
func (t *todoRepository) GetTodo(todo *domain.Todo, id string) (err error) {
	err = t.conn.Where("id = ?", id).First(todo).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodo update todo
func (t *todoRepository) UpdateTodo(todo *domain.Todo, id string) (err error) {
	err = t.conn.Save(todo).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo delete todo
func (t *todoRepository) DeleteTodo(todo *domain.Todo, id string) (err error) {
	err = t.conn.Where("id = ?", id).Delete(todo).Error
	if err != nil {
		return err
	}
	return nil
}
