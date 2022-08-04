package controllers

import (
	"net/http"

	"github.com/akposieyefa/crud-api/initializers"
	"github.com/akposieyefa/crud-api/models"
	"github.com/gin-gonic/gin"
)

//CREATE POST
func CreatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post created successfully",
			"success": true,
		})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to create post",
			"success": true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"success": true,
	})
}

//GET ALL POSTS
func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post fetched successfully",
		"posts":   posts,
		"success": true,
	})
}

//GET SINGLE POST
func GetSinglePost(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post fetched successfully",
		"posts":   post,
		"success": true,
	})
}

//UPDATE POST
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
		"success": true,
	})
}

//DELETE POST
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
		"success": true,
	})
}
