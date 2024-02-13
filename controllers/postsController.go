package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tymofiismyrnov/go_crud_api/initializers"
	"github.com/tymofiismyrnov/go_crud_api/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"post": post})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return posts
	c.IndentedJSON(http.StatusOK, gin.H{"posts": posts})
}

func PostsById(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Return posts
	c.IndentedJSON(http.StatusOK, gin.H{"post": post})
}

func PostsUpdate(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")
	// Get the data of the body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.IndentedJSON(http.StatusOK, post)
}

func PostDeleteById(c *gin.Context) {
	// Get ID from url
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)
	if post.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{fmt.Sprintf("id %v", id): "not found"})
		return
	}
	initializers.DB.Delete(&post, id)
	c.IndentedJSON(http.StatusOK, gin.H{"post_deleted": post})
}
