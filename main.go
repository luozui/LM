package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luozui/LM/db"
	"github.com/luozui/LM/server"
)

func main() {
	db.Init()
	r := gin.Default()
	r.GET("/", server.Index)
	r.POST("/add", server.Add)
	r.POST("/del", server.Del)

	r.POST("/add_docker", server.AddDocker)
	r.POST("/del_docker", server.OptDocker)
	r.POST("/top_docker", server.TopDocker)
	r.Run(":8080")
}
