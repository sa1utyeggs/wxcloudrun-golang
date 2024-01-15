package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wxcloudrun-golang/utils"
)

func CheckSource(c *gin.Context) {
	req := VerifyWechatRequest{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	// 验证消息来自 wechat 服务器
	if utils.VerifyInfoFromWechat("123123", req.Timestamp, req.Nonce, req.Signature) {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": "request illegal: not from wechat",
		})
		return
	}
}
