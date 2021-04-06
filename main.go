package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luozui/LM/db"
	"github.com/luozui/LM/server"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 可将将* 替换为指定的域名
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {
	db.Init()
	r := gin.Default()
	r.Use(Cors())
	r.StaticFile("/", "./static/index.html")
	r.GET("/get", server.Index)
	r.POST("/add", server.Add)
	r.POST("/opt", server.Opt)

	r.POST("/add_docker", server.AddDocker)
	r.POST("/opt_docker", server.OptDocker)
	r.POST("/top_docker", server.TopDocker)

	r.Run(":8080")
}
