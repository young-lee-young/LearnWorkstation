package binary

import (
	"testing"
	"fmt"
	"practice/utils/random"
	"practice/base_sort/bubble"
)

func TestBinarySearch(t *testing.T) {
	count := 10
	max := 20
	nums := random.GenerateArray(count, max)
	// 排序
	numsSorted := bubble.BubbleSort(nums)
	fmt.Println(numsSorted)
	target := random.GenerateIntFromMax(max)
	fmt.Println(target)
	result := binarySearch(numsSorted, target)
	fmt.Println(result)
}
