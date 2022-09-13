package selection

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestSelectionSorted(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := SelectionSort(nums)
	fmt.Println(numsSorted)
}
