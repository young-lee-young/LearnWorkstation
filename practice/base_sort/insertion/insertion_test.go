package insertion

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestInsertionSort(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := InsertionSort(nums)
	fmt.Println(numsSorted)
}
