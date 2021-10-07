package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", HeadReturn)

	// 4.当访问 localhost/healthz 时，应返回 200
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}

func HeadReturn(c *gin.Context) {
	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range c.Request.Header {
		for _, sv := range v {
			c.Writer.Header().Set(k, sv)
		}
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	c.Writer.Header().Set("VERSION", version)

	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	c.JSON(200, gin.H{})
}
