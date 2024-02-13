package main

import (
	"github.com/tymofiismyrnov/go_crud_api/initializers"
	"github.com/tymofiismyrnov/go_crud_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
