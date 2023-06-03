package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
)

// ToDoController use for handle framework here and present as a controller
type ToDoController struct {
	todoUseCase domains.ToDoUseCase
}

// NewToDoController ...
func NewToDoController(usecase domains.ToDoUseCase) *ToDoController {
	return &ToDoController{
		todoUseCase: usecase,
	}
}

// GetAllTodo ...
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
	var newToDo models.Todo
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

// GetTodo ...
func (t *ToDoController) GetTodo(c *gin.Context) {
	var newToDo models.Todo
	id := c.Params.ByName("id")
	err := t.todoUseCase.GetTodo(&newToDo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newToDo)
	}
}

// UpdateTodo ...
func (t *ToDoController) UpdateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var reqToDo models.Todo
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

// DeleteTodo ...
func (t *ToDoController) DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var respToDo models.Todo
	err := t.todoUseCase.DeleteTodo(&respToDo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
