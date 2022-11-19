package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 加密
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// Md5EncodeToUppercase 加密后转大写
func Md5EncodeToUppercase(str string) string {
	return strings.ToUpper(Md5Encode(str))
}
