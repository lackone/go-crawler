package utils

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// 获取编码
func GetEncode(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}
