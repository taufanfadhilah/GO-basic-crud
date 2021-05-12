package main

import (
	"go-basic-crud/handler"
	"go-basic-crud/task"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// define database
	dsn := "root:@tcp(127.0.0.1:3306)/go-basic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// to migrating db schema based on defined entities
	db.AutoMigrate(&task.Task{})

	// define repository, service, and handler
	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	api := router.Group("/api")

	// task routes
	api.GET("/task", taskHandler.Index)
	api.POST("/task", taskHandler.Store)
	api.GET("/task/:id", taskHandler.Show)
	api.PUT("/task/:id", taskHandler.Update)
	api.DELETE("/task/:id", taskHandler.Destory)

	router.Run()
}
