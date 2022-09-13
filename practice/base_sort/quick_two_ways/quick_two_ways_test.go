package quick_two_ways

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestQuickSortTwoWays(t *testing.T) {
	count := 50000
	max := 10
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := QuickSortTwoWays(nums)
	fmt.Println(numsSorted)
}
