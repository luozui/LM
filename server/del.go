package server

import (
	"github.com/gin-gonic/gin"
)

func Del(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
