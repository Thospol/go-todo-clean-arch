package repositories

import (
	_ "github.com/go-sql-driver/mysql" // use to connect db
	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

type todoRepository struct {
	conn *gorm.DB
}

// NewToDoRepository ...
func NewToDoRepository(conn *gorm.DB) domains.ToDoRepository {
	return &todoRepository{conn}
}

func (t *todoRepository) GetAllTodo(todo *[]models.Todo) (err error) {
	if err = t.conn.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) CreateTodo(todo *models.Todo) (err error) {
	if err = t.conn.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) GetTodo(todo *models.Todo, id string) (err error) {
	if err := t.conn.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateTodo(todo *models.Todo, id string) (err error) {
	// fmt.Println(todo)
	t.conn.Save(todo) // save all field
	return nil
}

func (t *todoRepository) DeleteTodo(todo *models.Todo, id string) (err error) {
	t.conn.Where("id = ?", id).Delete(todo)
	return nil
}
