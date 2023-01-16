package main

import "fmt"

/**
LeetCode No088 归并两个有序数组
*/
func main() {
	arrayOne := [9]int{1, 3, 5, 7, 9, 0, 0, 0, 0}
	arrayTwo := [4]int{2, 4, 6, 8}
	MergeTwoList(arrayOne, 5, arrayTwo, 4)
}

func MergeTwoList(arrayOne [9]int, m int, arrayTwo [4]int, n int) {
	p := m + n - 1
	n--
	m--

	for m >= 0 && n >= 0 {
		if arrayOne[m] >= arrayTwo[n] {
			arrayOne[p] = arrayOne[m]
			m--
		} else {
			arrayOne[p] = arrayTwo[n]
			n--
		}
		p--
	}

	for n >= 0 {
		arrayOne[p] = arrayTwo[n]
		n--
		p--
	}

	fmt.Println(arrayOne)
}
