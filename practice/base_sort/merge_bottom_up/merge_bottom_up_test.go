package merge_bottom_up

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestMergeSortBottomUp(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := MergeSortBottomUp(nums)
	fmt.Println(numsSorted)
}
