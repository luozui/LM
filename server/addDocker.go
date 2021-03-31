package server

import (
	"github.com/gin-gonic/gin"
)

func AddDocker(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
