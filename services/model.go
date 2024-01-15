package services

type VerifyWechatRequest struct {
	Signature string `form:"signature"` // 微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。
	Timestamp string `form:"timestamp"` // 时间戳
	Nonce     string `form:"nonce"`     // 随机数
	EchoStr   string `form:"echostr"`   // 随机字符串
}

type EventRequest struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	OpenID       string `json:"OpenID"`
	AppID        string `json:"AppID"`
	RevokeInfo   string `json:"RevokeInfo"`
	PluginID     string `json:"PluginID"`
	OpenPID      string `json:"OpenPID"`
}
