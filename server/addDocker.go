package server

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddDocker(c *gin.Context) {
	dockerName := c.DefaultPostForm("dockername", "")
	dockerTag := c.DefaultPostForm("dockertag", "")
	cpus := c.DefaultPostForm("cpus", "16")
	mem := c.DefaultPostForm("mem", "32G")
	gpus := c.DefaultPostForm("gpus", "")
	ip := c.DefaultPostForm("ip", "")
	homePath := c.DefaultPostForm("homepath", "")

	if dockerName == "" {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "dockerName is nil",
		})
		return
	}

	cmdargs := fmt.Sprintf("run -d --restart=always --net=mcv --ip=%v --gpus=%v --cpus %v -m %v -v /%v:/notebooks --name lm_%v -e PASSWORD=\"gzdx\" -e PORT=\"80\" %v", ip, gpus, cpus, mem, homePath, dockerName, dockerTag)
	agrs := strings.Fields(strings.TrimSpace(cmdargs))
	cmd := exec.Command("docker", agrs...)
	log.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err == nil {
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
