package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tymofiismyrnov/go_crud_api/initializers"
)

func RedisSet(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.String(http.StatusBadRequest, "'key' param is not found")
		return
	}
	value := c.Query("value")
	if value == "" {
		c.String(http.StatusBadRequest, "'value' param is not found")
		return
	}
	err := initializers.RDB.Set(initializers.CTX, key, value, 0).Err()
	if err != nil {
		c.String(http.StatusBadGateway, fmt.Sprintf("%v", err))
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "key-value pair set successfuly",
		"key":     key,
		"value":   value,
	})

}

func RedisGet(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.String(http.StatusBadRequest, "'key' param is not found")
		return
	}
	val, err := initializers.RDB.Get(initializers.CTX, key).Result()
	if err == redis.Nil {
		c.String(http.StatusNotFound, fmt.Sprintf("%v does not exist", key))
	} else if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("get %v value successfuly", key),
			"value":   val,
		})
	}
}

func RedisDelete(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.String(http.StatusBadRequest, "'key' param is not found")
		return
	}
	val, err := initializers.RDB.Del(initializers.CTX, key).Result()
	if err == redis.Nil {
		c.String(http.StatusNotFound, fmt.Sprintf("%v does not exist", key))
	} else if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfuly deleted %v value", key),
			"key":     key,
			"value":   val,
		})
	}

}
