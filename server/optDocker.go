package server

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func OptDocker(c *gin.Context) {
	dockerName := c.DefaultQuery("dockername", "")
	opt := c.DefaultQuery("dockername", "stop")
	if dockerName == "" {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "dockerName is nil",
		})
		return
	}
	cmd := exec.Command("docker", opt, dockerName)
	log.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     string(out),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": err == nil,
		"msg":     fmt.Sprintf(err.Error()),
	})
}
