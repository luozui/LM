package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luozui/LM/db"
	"github.com/luozui/LM/model"
)

func Add(c *gin.Context) {
	dockerName := c.DefaultPostForm("dockername", "")
	dockerTag := c.DefaultPostForm("dockertag", "")
	cpus := c.DefaultPostForm("cpus", "16")
	mem := c.DefaultPostForm("mem", "32")
	gpus := c.DefaultPostForm("gpus", "0")
	ip := getIP()
	machineip := c.DefaultPostForm("machineip", "")
	homePath := c.DefaultPostForm("homepath", "")
	endtime := c.DefaultPostForm("endtime", "4")

	names := c.DefaultPostForm("names", "")
	description := c.DefaultPostForm("description", "")
	password := c.DefaultPostForm("password", "")

	urlValues := url.Values{
		"dockername": {dockerName},
		"dockertag":  {dockerTag},
		"cpus":       {cpus},
		"gpus":       {gpus},
		"mem":        {mem},
		"ip":         {ip},
		"Machineip":  {machineip},
		"homepath":   {homePath},
		"token":      {db.Data.Token},
		"password":   {password},
	}
	// urlValues.Add("name", "zhaofan")
	// urlValues.Add("age", "22")
	resp, _ := http.PostForm(fmt.Sprintf("http://%v:8080/add_docker", machineip), urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	var resp_ gin.H
	err := json.Unmarshal(body, &resp_)
	if err != nil {
		log.Println("error:", err)
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "resp error",
		})
		db.Data.IP[ip] = false
		return
	}
	adduser(dockerName, dockerTag, cpus, gpus, mem, ip, machineip, homePath, endtime, names, description, string(md5.New().Sum([]byte(password))))
	addLoad(machineip, cpus, gpus, mem)
	db.Write()
	c.JSON(200, resp_)
}

func adduser(dockerName, dockerTag, cpus, gpus, mem, ip, machineip, homePath, endtime, names, description, password string) {
	var et int64
	fmt.Sscanf(endtime, "%d", &et)
	if db.Data.Users[dockerName] == nil {
		db.Data.Users[dockerName] = &model.User{
			DockerName:  dockerName,
			DockerTag:   dockerTag,
			Names:       names,
			Description: description,
			IP:          ip,
			Machineip:   machineip,
			CPUs:        cpus,
			GPUs:        gpus,
			Mem:         mem,
			StartTime:   time.Now(),
			Status:      1,
			Password:    password,
			EndTime:     time.Now().Add(time.Duration(et * 1000000000 * 60)), // 小时
		}
	}
}

func addLoad(ip, cpus, gpus, mem string) {
	intgpus := [4]int{-1, -1, -1, -1}
	var intcpus int
	intmem := 16
	fmt.Sscanf(cpus, "%d", &intcpus)
	fmt.Sscanf(gpus, "%d,%d,%d,%d", &intgpus[0], &intgpus[1], &intgpus[2], &intgpus[3])
	fmt.Sscanf(mem, "%d", &intmem)
	num := 0
	for _, i := range intgpus {
		if i > 0 {
			num++
			db.Data.Machines[ip].Use[i]++
		}
	}
	db.Data.Machines[ip].Load[1] += num
	db.Data.Machines[ip].Load[2] += intcpus
	db.Data.Machines[ip].Load[3] += intmem
}

func delLoad(ip, cpus, gpus, mem string) {
	intgpus := [4]int{0, 0, 0, 0}
	var intcpus int
	intmem := 16
	fmt.Sscanf(cpus, "%d", &intcpus)
	fmt.Sscanf(gpus, "%d,%d,%d,%d", &intgpus[0], &intgpus[1], &intgpus[2], &intgpus[3])
	fmt.Sscanf(mem, "%d", &intmem)
	num := 0
	for _, i := range intgpus {
		if i > 0 {
			num++
			db.Data.Machines[ip].Use[i]--
		}
	}
	db.Data.Machines[ip].Load[1] -= num
	db.Data.Machines[ip].Load[2] -= intcpus
	db.Data.Machines[ip].Load[3] -= intmem
}

func getMinCpuLoadMachine() string {
	ip := ""
	load := 0x3f3f3f3f
	for k, m := range db.Data.Machines {
		if m.Load[2] < load {
			load = m.Load[1]
			ip = k
		}
	}
	return ip
}

func getMinGpuLoadMachine() string {
	ip := ""
	load := 0x3f3f3f3f
	for k, m := range db.Data.Machines {
		if m.Load[1] < load {
			load = m.Load[1]
			ip = k
		}
	}
	return ip
}

func getIP() string {
	for ip, use := range db.Data.IP {
		if !use {
			db.Data.IP[ip] = true
			return ip
		}
	}
	return "error"
}
