package print

import "fmt"

/**
	打印一维数组
 */
func PrintOneDimensionalArray(numArray []int) {
	for i := 0; i < len(numArray); i ++ {
		fmt.Println(numArray[i])
	}
	fmt.Println("------------------------")
}

/**
	打印二维数组
 */
func PrintTwoDimensionalArray(numArray [][]int) {
	for i := 0; i < len(numArray); i ++ {
		fmt.Println(numArray[i])
	}
	fmt.Println("------------------------")
}
