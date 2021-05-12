package main

import (
	"go-basic-crud/handler"
	"go-basic-crud/todo"
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
	db.AutoMigrate(&todo.Todo{})

	todoRepository := todo.NewRepository(db)
	todoService := todo.NewService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// todo routes
	api.POST("/todo", todoHandler.Store)

	router.Run()
}
