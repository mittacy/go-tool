package encodeUtil

import (
	"crypto/md5"
	"encoding/hex"
)

// Get16MD5Encode 返回16位md5加密后的字符串
func Get16MD5Encode(str string) string {
	return Get32MD5Encode(str)[8:24]
}

// Get32MD5Encode 返回32位md5加密后的字符串
func Get32MD5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
