/**
* @Time:    2023/3/2 10:40
* @Author:  leeyoung
* @File:    main.go
* @Content:

时针 分针 夹角

参数：

小时：

分钟：
*/

package main

import "fmt"

func main() {
	ret := solution(3, 50)

	fmt.Println("ret", ret)
}

/**
小时：360 / 12 = 30度

分钟：360 / 60 = 6 度

每走 1分钟 时针走的角度：30 / 60 = 0.5
*/
func solution(hour float64, minute float64) float64 {
	// 1:30
	// hour = 1，minute = 30

	// 时针相对于 0 点角度
	i := hour*30 + minute*0.5

	// 分针相对于 0 点角度
	j := minute * 6

	ret := j - i

	return ret
}
