package main

import (
	"fmt"
	"os"
)

/**
	文件名任意，包名和方法名必须是main
 */
func main() {
	// main方法没有参数，用os.Args获取参数
	fmt.Println(os.Args[1])
}
