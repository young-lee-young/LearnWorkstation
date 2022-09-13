package shell

import (
	"testing"
	"practice/utils/random"
	"fmt"
)

func TestShellSort(t *testing.T) {
	count := 50000
	max := 1000000
	nums := random.GenerateArray(count, max)
	fmt.Println(nums)

	numsSorted := ShellSort(nums)
	fmt.Println(numsSorted)
}
