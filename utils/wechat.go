package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

func VerifyInfoFromWechat(token, timestamp, nonce, signature string) bool {
	data := []string{token, timestamp, nonce}
	sort.Strings(data)
	// 创建SHA-1哈希对象
	hasher := sha1.New()
	// 将字符串输入写入哈希对象
	hasher.Write([]byte(strings.Join(data, "")))
	// 计算哈希值
	hashBytes := hasher.Sum(nil)
	// 将哈希值转换为16进制字符串
	return hex.EncodeToString(hashBytes) == signature
}
