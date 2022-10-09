package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "498828660196"
	num2 := "840477629533"
	//num1 := "2"
	//num2 := "3"
	ret := solution(num1, num2)
	fmt.Println("ret:", ret)

	//fmt.Println("8:", 8 * 498828660196)
	//fmt.Println("4:", 4 * 498828660196)
	//8: 3990629281568
	//4: 1995314640784
}

func solution(num1 string, num2 string) string {
	ret := ""
	int1, _ := strconv.Atoi(num1)

	for index, num := range num2 {
		numInt, _ := strconv.Atoi(string(num))
		// 位上的数字和 num1 相乘
		temp := strconv.Itoa(numInt * int1)

		i := 0
		for i < len(num2)-index {
			temp += "0"
			i++
		}
		fmt.Println("temp", temp)

		cur, _ := strconv.ParseInt(temp, 10, 64)

		ret += cur
		fmt.Println("ret", ret)
	}
}
