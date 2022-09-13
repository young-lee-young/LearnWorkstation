package quick_three_ways

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestQuickSortThreeWays(t *testing.T) {
	count := 50000
	max := 10
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := QuickSortThreeWays(nums)
	fmt.Println(numsSorted)
}
