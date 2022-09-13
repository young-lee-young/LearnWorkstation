package bubble

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := BubbleSort(nums)
	fmt.Println(numsSorted)
}
