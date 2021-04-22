package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

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
	password := c.DefaultPostForm("password", "")

	if dockerName == "" {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "dockerName is nil",
		})
		return
	}

	filedata := fmt.Sprintf("#!/bin/bash\n nvidia-docker run -d --restart=always --net=mcv --ip=%v --gpus=%v --cpus=%v -m %v -v /lm_data/%v:/notebooks --name lm_%v -e PASSWORD=\"%v\" -e PORT=\"80\" %v", ip, gpus, cpus, mem, homePath, dockerName, password, dockerTag)
	err := ioutil.WriteFile("run.sh", []byte(filedata), 0655)
	if err != nil {
		log.Println("load file error")
		c.JSON(200, gin.H{
			"success": false,
			"msg":     fmt.Sprintf(err.Error()),
		})
		return
	}
	// cmdargs := fmt.Sprintf("#!/bin/bash\n nvidia-docker run -d --restart=always --net=mcv --ip=%v --gpus=%v --cpus=%v -m %v -v /lm_data/%v:/notebooks --name lm_%v -e PASSWORD=\"gzdx\" -e PORT=\"80\" %v", ip, gpus, cpus, mem, homePath, dockerName, dockerTag)
	// agrs := strings.Fields(strings.TrimSpace(cmdargs))
	cmd := exec.Command("./run.sh")
	log.Println(cmd)
	err = cmd.Run()
	if err == nil {
		c.JSON(200, gin.H{
			"success": true,
			"ip":      ip,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": err == nil,
		"msg":     fmt.Sprintf(err.Error()),
	})
}
