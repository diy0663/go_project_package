package helper

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"time"
)

// 常用函数

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// 指定长度的数字验证码
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < length; i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// 指定长度的随机字符串
func RandomString(length int) string {

	mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}
