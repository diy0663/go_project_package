package helper

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

// 常用函数

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

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
