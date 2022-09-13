package main

import (
	"fmt"
	"os"
)

func main() {
	// println输出到标注错误，没有返回值
	println("hello world")

	// fmt.Println输出到标准输出，有返回值：1. 写入到标准输出字节数 2.发生的错误
	num, err := fmt.Println("hello world")
	if err != nil {
		fmt.Println("error happened")
	} else {
		fmt.Printf("%v bytes have be written in stdout", num)
	}

	// 变量
	write, err := fmt.Printf("hello %s\n", "李")
	fmt.Println(write, err)

	ret := fmt.Sprintf("hello %s", "李")
	fmt.Println("ret", ret)

	n, err := fmt.Fprintf(os.Stdout, "hello %s\n", "李")
	fmt.Println(n, err)
}
