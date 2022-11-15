package config

import (
	"fmt"
	"skyshi-rest-api/controllers"
	"skyshi-rest-api/models"
	"skyshi-rest-api/repositories"
	"skyshi-rest-api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	fmt.Println("Welcome my rest api")

	//migrate
	db.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)

	PORT := ":3030"
	router := echo.New()

	// Activity
	activityRepository := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepository)
	activityController := controllers.NewActivityController(activityService)
	router.GET("/activity-groups", activityController.GetAllActivity)
	router.GET("/activity-groups/:id", activityController.GetActivity)
	router.POST("/activity-groups", activityController.CreateActivity)
	router.PATCH("/activity-groups/:id", activityController.UpdateActivity)
	router.DELETE("/activity-groups/:id", activityController.DeleteActivity)

	// Todo
	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)
	router.GET("/todo-items", todoController.GetAllTodo)
	router.GET("/todo-items/:id", todoController.GetTodo)
	router.POST("/todo-items", todoController.CreateTodo)
	router.PATCH("/todo-items/:id", todoController.UpdateTodo)
	router.DELETE("/todo-items/:id", todoController.DeleteTodo)

	router.Logger.Fatal(router.Start(PORT))
}
