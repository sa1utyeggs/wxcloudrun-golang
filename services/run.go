package services

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xlstudio/wxbizdatacrypt"
	"io/ioutil"
	"net/http"
)

func requestLoggerMiddleware(c *gin.Context) {
	// 打印请求URI
	fmt.Printf("Request URI: %s\n", c.Request.RequestURI)

	// 读取请求体
	body, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		// 打印请求体
		fmt.Printf("Request Body: %s\n", body)

		// 重置请求体，确保后续处理中间件和路由处理函数可以正常读取
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	// 继续处理请求
	c.Next()
}

func Run() {
	server := gin.Default()

	// Get 请求 返回 JSON
	applet := server.Group("/wechat-server", requestLoggerMiddleware)
	{
		applet.POST("/revoke", func(c *gin.Context) {
			params := VerifyWechatRequest{}
			err := c.ShouldBindQuery(&params)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			request := RevokeEventRequest{}
			err = c.ShouldBindJSON(&request)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			// 解密
			pc := wxbizdatacrypt.WxBizDataCrypt{AppId: request.AppID, SessionKey: request.SessionKey}
			_, err = pc.Decrypt(request.Encrypt, params.Nonce, true)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			// TODO 记录、解绑

			c.String(http.StatusOK, Success)
		})
	}

	// 指定运行在 8080 端口
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
