package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luozui/LM/db"
)

func Index(c *gin.Context) {
	//resp, _ := json.Marshal(db.Data)
	c.JSON(200, gin.H{
		"success":  true,
		"uses":     db.Data.Users,
		"machines": db.Data.Machines,
	})
}
