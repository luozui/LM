package server

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func AddDocker(c *gin.Context) {
	dockerName := c.DefaultQuery("dockername", "")
	dockerTag := c.DefaultQuery("dockertag", "")
	cpus := c.DefaultQuery("cpus", "16")
	mem := c.DefaultQuery("mem", "32G")
	gpus := c.DefaultQuery("gpus", "")
	ip := c.DefaultQuery("gpus", "")
	homePath := c.DefaultQuery("homepath", "")

	if dockerName == "" {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "dockerName is nil",
		})
		return
	}

	cmdargs := fmt.Sprintf("-d --net=mcv --ip=%v --gpus=%v --cpus %v -m %v -v /%v:/mydata --name %v %v", ip, gpus, cpus, mem, homePath, dockerName, dockerTag)
	cmd := exec.Command("docker", "run", cmdargs)
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
