package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xlstudio/wxbizdatacrypt"
	"net/http"
)

func Run() {
	server := gin.Default()

	// Get 请求 返回 JSON
	applet := server.Group("/wechat-server", CheckSource)
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
	err := server.Run(":8081")
	if err != nil {
		return
	}
}
