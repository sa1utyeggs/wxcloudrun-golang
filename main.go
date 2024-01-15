package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wxcloudrun-golang/services"
	"wxcloudrun-golang/utils"
)

func main() {
	server := gin.Default()

	server.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

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
