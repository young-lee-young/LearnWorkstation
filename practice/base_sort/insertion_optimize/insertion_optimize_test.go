package insertion_optimize

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestInsertionOptimize(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := InsertionSortOptimize(nums)
	fmt.Println(numsSorted)
}
