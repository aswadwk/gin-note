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

	activityRepo repositories.ActivityRepository = repoImpl.CreateActivityRepository(db)

	activityService services.ActivityService = serviceImpl.CreateActivityService(activityRepo)

	activityController controllers.ActivityController = controllers.NewActivityController(activityService)
)

func main() {
	// ...
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	r := setupRouter()

	r.POST("/activity-groups", activityController.Create)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	r.Run()

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}