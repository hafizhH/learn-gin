package main

import (
	"LearnAPI/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookById)
	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:id", controllers.UpdateBookById)
	router.DELETE("/books/:id", controllers.DeleteBookById)

	router.Run("localhost:8080")
}
