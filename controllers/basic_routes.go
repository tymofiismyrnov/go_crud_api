package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Favicon(c *gin.Context) {
	filepath := "./public/favicon.png"
	c.File(filepath)
}

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
