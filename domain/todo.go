package domain

// Todo ...
type Todo struct {
	Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TableName use to specific table
func (b *Todo) TableName() string {
	return "todo"
}

// ToDoUseCase ...
type ToDoUseCase interface {
	GetAllTodo() (t []Todo, err error)
	CreateTodo(t *Todo) (err error)
	GetTodo(t *Todo, id string) (err error)
	UpdateTodo(t *Todo, id string) (err error)
	DeleteTodo(t *Todo, id string) (err error)
}

// ToDoRepository ...
type ToDoRepository interface {
	GetAllTodo(t *[]Todo) (err error)
	CreateTodo(t *Todo) (err error)
	GetTodo(t *Todo, id string) (err error)
	UpdateTodo(t *Todo, id string) (err error)
	DeleteTodo(t *Todo, id string) (err error)
}
