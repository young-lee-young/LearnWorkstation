package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var username string

	username = "leeyoung"
	fmt.Println(username)
	fmt.Println(len(username))

	username = "\xE4\xB8\xA5"
	fmt.Println(username)
	fmt.Println(len(username))

	stringProcess()
	unicodeLearn()
}

// 字符串操作
func stringProcess() {
	// 字符串分割
	testStr := "A,B,C"
	strArray := strings.Split(testStr, ",")
	for index, value := range strArray {
		fmt.Println(index, value)
	}

	// 字符串拼接
	testStrNew := strings.Join(strArray, "-")
	fmt.Println(testStrNew)

	// 整形转字符串
	intToStr := strconv.Itoa(10)
	fmt.Println(intToStr + "元")

	// 字符串转整形
	if value, err := strconv.Atoi("10"); err == nil {
		fmt.Println(value + 1)
	}
}

// Unicode和UTF-8
func unicodeLearn()  {
	username := "李二狗"

	fmt.Println(username)
	// 是byte数
	fmt.Println(len(username))

	usernameRune := []rune(username)
	fmt.Println(len(usernameRune))
}
