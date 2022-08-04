package main

import (
	"github.com/akposieyefa/crud-api/controllers"
	"github.com/akposieyefa/crud-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
