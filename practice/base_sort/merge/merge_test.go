package merge

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestMergeSort(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := MergeSort(nums)
	fmt.Println(numsSorted)
}
