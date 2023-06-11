package route

import (
	controllers "github.com/Thospol/go-todo-clean-arch/api/controller"
	"github.com/Thospol/go-todo-clean-arch/database"
	"github.com/Thospol/go-todo-clean-arch/repository"
	"github.com/Thospol/go-todo-clean-arch/usecase"
	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {
	todoRepo := repository.NewToDoRepository(database.DB)
	todoUseCase := usecase.NewToDoUseCase(todoRepo)
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
