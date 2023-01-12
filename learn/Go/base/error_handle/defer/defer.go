package main

import (
	"bufio"
	"errors"
	"fmt"
	"functional/fib"
	"os"
)

/**
defer确保调用在函数结束时发生
参数在defer语句时计算
defer列表为后进先出
*/
func tryDefer() {
	// defer 放入的数据是一个栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func paramDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 8 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func openFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 自定义error
	err = errors.New("this is a custom error")

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err.Error())
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	} else {
		file.Close()
	}
}

func main() {
	const filename = "/Users/leeyoung/WorkStation/LearnWorkstation/learn/Go/src/error_handle/defer/fib.txt"
	writeFile(filename)
	openFile(filename)
	paramDefer()
	tryDefer()
}
