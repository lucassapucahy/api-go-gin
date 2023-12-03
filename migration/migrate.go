package main

import (
	initializers "GoGinApi/Initializers"
	"GoGinApi/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}
