package randomUtil

import (
	"math/rand"
	"time"
)

var (
	defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// String 生成随机字符串
// @param n 生成的字符串长度
// @return string 返回生成的随机字符串
func String(n int, randChars ...[]rune) string {
	if n <= 0 {
		return ""
	}

	var letters []rune

	if len(randChars) == 0 {
		letters = defaultLetters
	} else {
		letters = randChars[0]
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// Num 使用math/rand库的随机数，伪随机数，不建议使用在密码类
// @param min 范围最小
// @param max 范围最大
// @return int64
func Num(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
