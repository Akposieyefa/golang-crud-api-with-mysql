package main

import (
	"github.com/akposieyefa/crud-api/initializers"
	"github.com/akposieyefa/crud-api/models"
)

func init() {
	initializers.ConectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
