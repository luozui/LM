package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luozui/LM/db"
	"github.com/luozui/LM/server"
)

func main() {
	db.Init()
	r := gin.Default()
	r.StaticFile("/", "./static/index.html")
	r.GET("/get", server.Index)
	r.POST("/add", server.Add)
	r.POST("/opt", server.Opt)

	r.POST("/add_docker", server.AddDocker)
	r.POST("/opt_docker", server.OptDocker)
	r.POST("/top_docker", server.TopDocker)

	r.Run(":8080")
}
