package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tymofiismyrnov/go_crud_api/controllers"
	"github.com/tymofiismyrnov/go_crud_api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Basic routes
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/favicon.ico", controllers.Favicon)
	r.GET("/healthz", controllers.HealthCheck)

	// Posts routes
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsById)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDeleteById)

	// Run server
	r.Run()
}
