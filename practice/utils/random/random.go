package random

import (
	"math/rand"
	"time"
)

// 生成随机数组
func GenerateArray(count int, max int) []int {
	numSlice := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		numSlice = append(numSlice, rand.Intn(max))
	}
	return numSlice
}

/**
给定最大值生成随机数
*/
func GenerateIntFromMax(max int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max)
	return num
}

/**
给定最大最小值生成随机数
*/
func GenerateIntFromMinMax(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	num := min + rand.Intn(max-min)
	return num
}
