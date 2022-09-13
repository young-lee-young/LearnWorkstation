package quick

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := QuickSort(nums)
	fmt.Println(numsSorted)
}
