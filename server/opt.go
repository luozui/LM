package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/luozui/LM/db"
)

func Opt(c *gin.Context) {
	dockerName := c.DefaultPostForm("dockername", "")
	machineip := db.Data.Users[dockerName].Machineip
	opt := c.DefaultPostForm("opt", "")
	names := c.DefaultPostForm("names", db.Data.Users[dockerName].Names)
	description := c.DefaultPostForm("description", "")
	password := c.DefaultPostForm("password", "")

	urlValues := url.Values{
		"dockername": {dockerName},
		"opt":        {opt},
		"token":      {db.Data.Token},
	}

	if string(md5.New().Sum([]byte(password))) != db.Data.Users[dockerName].Password {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "passowrd error",
		})
		return
	}

	log.Println(dockerName, machineip, opt, names, description)
	var resp_ gin.H
	if opt != "" {
		resp, _ := http.PostForm(fmt.Sprintf("http://%v:8080/opt_docker", machineip), urlValues)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("error:", err)
			c.JSON(200, gin.H{
				"success": false,
				"msg":     "resp error",
			})
			return
		}

		err = json.Unmarshal(body, &resp_)
		if err != nil {
			log.Println("error:", err)
			c.JSON(200, gin.H{
				"success": false,
				"msg":     "resp error",
			})
			return
		}
		log.Println(string(body))
	}

	// 更新 user 数据
	if names != "" {
		db.Data.Users[dockerName].Names = names
	}
	if description != "" {
		db.Data.Users[dockerName].Description = description
	}
	switch opt {
	case "start":
		if db.Data.Users[dockerName].Status != 1 {
			addLoad(machineip, db.Data.Users[dockerName].CPUs, db.Data.Users[dockerName].GPUs, db.Data.Users[dockerName].Mem)
		}
		db.Data.Users[dockerName].Status = 1
		break
	case "stop":
		if db.Data.Users[dockerName].Status == 1 {
			delLoad(machineip, db.Data.Users[dockerName].CPUs, db.Data.Users[dockerName].GPUs, db.Data.Users[dockerName].Mem)
		}
		db.Data.Users[dockerName].Status = 2
		break
	case "rm":
		if db.Data.Users[dockerName].Status == 2 {
			//delLoad(machineip, db.Data.Users[dockerName].CPUs, db.Data.Users[dockerName].GPUs, db.Data.Users[dockerName].Mem)
			delete(db.Data.Users, dockerName)
		}
		break
	}
	db.Write()
	c.JSON(200, resp_)
}
