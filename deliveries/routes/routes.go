package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-clean-arch/database"
	"github.com/krittawatcode/go-todo-clean-arch/deliveries/controllers"
	"github.com/krittawatcode/go-todo-clean-arch/repositories"
	"github.com/krittawatcode/go-todo-clean-arch/usecases"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {
	todoRepo := repositories.NewToDoRepository(database.DB)
	todoUseCase := usecases.NewToDoUseCase(todoRepo)
	todoController := controllers.NewToDoController(todoUseCase)

	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("todo", todoController.GetAllTodo)
		v1.POST("todo", todoController.CreateTodo)
		v1.GET("todo/:id", todoController.GetTodo)
		v1.PUT("todo/:id", todoController.UpdateTodo)
		v1.DELETE("todo/:id", todoController.DeleteTodo)
	}
	return r
}
