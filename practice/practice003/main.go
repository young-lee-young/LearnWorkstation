package main

import (
	"fmt"
	"time"
)

var hashMap map[uint64]uint64

/**
	斐波那契数列 LeetCode No509

	标签：动态规划
 */
func main() {
	startTime := time.Now().Second()
	// 指数级增长，使用哈希表进行缓存
	hashMap = make(map[uint64]uint64)
	result := solution(100)
	fmt.Println(result)
	endTime := time.Now().Second()
	fmt.Println(endTime - startTime)
}

func solution(n uint64) uint64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		var sum1 uint64
		if val, ok := hashMap[n - 1]; ok {
			sum1 = val
		} else {
			sum1 = solution(n - 1)
		}
		var sum2 uint64
		if val, ok := hashMap[n - 2]; ok {
			sum2 = val
		} else {
			sum2 = solution(n - 2)
		}
		sum := sum1 + sum2
		hashMap[n] = sum
		return sum
	}
}
