package main

import (
	initializers "GoGinApi/Initializers"
	"GoGinApi/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/book", controllers.GetBooks)
	r.GET("/book/:id", controllers.GetBooksById)
	r.POST("/book", controllers.BookCreate)
	r.PUT("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.Deletebook)

	r.Run()
}
