package main

import (
	"math"
	"fmt"
)

/**
	两个数的平方和 LeetCode No633

	标签：双指针
 */
func main()  {
	sum := 2

	i := 0
	j := int(math.Ceil(math.Sqrt(float64(sum))))
	fmt.Println(j)

	for i <= j {
		sumTemp := i * i + j * j
		if sumTemp == sum {
			fmt.Println(i, j)
			return
		}
		if sumTemp > sum {
			j --
		} else {
			i ++
		}
	}
	fmt.Println(false)
}
