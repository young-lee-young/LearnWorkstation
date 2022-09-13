package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// UTF-8 每个中文占3个字节
	s := "yes我爱慕课网!"

	// len获取字节长度
	fmt.Println(len(s))

	// []byte获取字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}

	// 获取字符数
	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
