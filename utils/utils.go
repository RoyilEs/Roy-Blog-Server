package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// InList 是否存在列表里面
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Md5 加密
func Md5(str []byte) string {
	m := md5.New()
	m.Write(str)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
