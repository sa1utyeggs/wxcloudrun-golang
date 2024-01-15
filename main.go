package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wxcloudrun-golang/services"
	"wxcloudrun-golang/utils"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求时间
		// start := time.Now()

		// 打印请求信息
		reqBody, _ := c.GetRawData()
		fmt.Printf("[INFO] Request: %s %s %s\n", c.Request.Method, c.Request.RequestURI, reqBody)

		// 执行请求处理程序和其他中间件函数
		c.Next()

		// 记录回包内容和处理时间
		//end := time.Now()
		//latency := end.Sub(start)
		//respBody := string(c.Writer.Body.Bytes())
		//fmt.Printf("[INFO] Response: %s %s %s (%v)\n", c.Request.Method, c.Request.RequestURI, respBody, latency)
	}
}

func main() {
	server := gin.Default()

	server.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	}, Logger())

	server.GET("/verify-wechat", func(c *gin.Context) {
		req := services.VerifyWechatRequest{}
		err := c.ShouldBindQuery(&req)
		if err != nil {
			c.String(http.StatusBadRequest, "")
			return
		}

		// 验证消息来自 wechat 服务器
		if utils.VerifyInfoFromWechat("token", req.Timestamp, req.Nonce, req.Signature) {
			c.String(http.StatusOK, req.EchoStr)
			return
		}
		c.String(http.StatusBadRequest, "")
	})

	// Get 请求 返回 JSON
	applet := server.Group("/wechat-server")
	{
		applet.POST("/revoke", func(c *gin.Context) {
			request := services.EventRequest{}
			err := c.ShouldBindJSON(&request)
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusBadRequest, "")
				return
			}

			// TODO 记录、解绑

			c.String(http.StatusOK, services.Success)
		})
	}

	// 指定运行在 8080 端口
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
