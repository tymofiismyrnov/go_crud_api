package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	filepath := "./index.html"
	c.File(filepath)
}

func Favicon(c *gin.Context) {
	filepath := "./static/favicon.png"
	c.File(filepath)
}

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
