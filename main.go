package main

import (
	"github.com/gin-gonic/gin"

	"github.com/tymofiismyrnov/go_crud_api/controllers"
	"github.com/tymofiismyrnov/go_crud_api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectoToRedis()
}

func main() {
	r := gin.Default()

	// Basic routes
	r.Static("/static", "./static")
	r.GET("/", controllers.Index)
	r.GET("/favicon.ico", controllers.Favicon)
	r.GET("/healthz", controllers.HealthCheck)

	// Posts routes
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsById)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostDeleteById)

	r.POST("/redis", controllers.RedisSet)
	r.GET("/redis", controllers.RedisGet)
	r.DELETE("/redis", controllers.RedisDelete)

	// Run server
	r.Run()
}
