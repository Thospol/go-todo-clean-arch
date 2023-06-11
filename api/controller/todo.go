package controller

import (
	"net/http"

	"github.com/Thospol/go-todo-clean-arch/domain"
	"github.com/gin-gonic/gin"
)

// ToDoController use for handle framework here and present as a controller
type ToDoController struct {
	todoUseCase domain.ToDoUseCase
}

// NewToDoController new instance of todo controller
func NewToDoController(usecase domain.ToDoUseCase) *ToDoController {
	return &ToDoController{
		todoUseCase: usecase,
	}
}

// GetAllTodo get all todo
func (t *ToDoController) GetAllTodo(c *gin.Context) {
	resp, err := t.todoUseCase.GetAllTodo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

// CreateTodo ...
func (t *ToDoController) CreateTodo(c *gin.Context) {
	var newToDo domain.Todo
	if err := c.ShouldBind(&newToDo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := t.todoUseCase.CreateTodo(&newToDo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newToDo)
	}
}

// GetTodo get todo
func (t *ToDoController) GetTodo(c *gin.Context) {
	var newToDo domain.Todo
	id := c.Params.ByName("id")
	err := t.todoUseCase.GetTodo(&newToDo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newToDo)
	}
}

// UpdateTodo update todo
func (t *ToDoController) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var reqToDo domain.Todo
	if err := c.ShouldBind(&reqToDo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := t.todoUseCase.UpdateTodo(&reqToDo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reqToDo)
	}
}

// DeleteTodo delete todo
func (t *ToDoController) DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var respToDo domain.Todo
	err := t.todoUseCase.DeleteTodo(&respToDo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
