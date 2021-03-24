package main

import (
	"github.com/gin-gonic/gin"
	"github.com/julian-mp/bookstore-api/controllers"
	"github.com/julian-mp/bookstore-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.GET("/books", controllers.FindBooks)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}