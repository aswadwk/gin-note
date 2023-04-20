package main

import (
	"aswadwk/config"
	"aswadwk/controllers"
	"aswadwk/repositories"
	repoImpl "aswadwk/repositories/impl"
	"aswadwk/services"
	serviceImpl "aswadwk/services/impl"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.DBConnect()

	// repo
	activityRepo repositories.ActivityRepository = repoImpl.CreateActivityRepository(db)
	todoRepo repositories.TodoRepository = repoImpl.CreateTodoRepository(db)

	// service
	activityService services.ActivityService = serviceImpl.CreateActivityService(activityRepo)
	todoService services.TodoService = serviceImpl.CreateTodoService(todoRepo)

	// controller
	activityController controllers.ActivityController = controllers.NewActivityController(activityService)
	todoController controllers.TodoController = controllers.NewTodoController(todoService)
)

func main() {
	// ...
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file main")
	}

	route := setupRouter()

	route.POST("/activity-groups", activityController.Create)
	route.GET("/activity-groups", activityController.FindAll)
	route.GET("/activity-groups/:id", activityController.FindByID)
	route.PATCH("/activity-groups/:id", activityController.UpdateByID)
	route.DELETE("/activity-groups/:id", activityController.DeleteByID)

	route.POST("/todo-items", todoController.Create)
	route.GET("/todo-items", todoController.FindAll)
	route.GET("/todo-items/:id", todoController.FindByID)
	route.PATCH("/todo-items/:id", todoController.UpdateByID)
	route.DELETE("/todo-items/:id", todoController.DeleteByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030" // Default port if not specified
	}
	route.Run()

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}