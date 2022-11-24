/**
 * @Time:    2022/11/24 18:38 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:
	🥇标准实现

	快速排序
 */
package main

import "fmt"

func main() {
	array := []int{19, 97, 9, 17, 1, 8}
	QuickSort(array, 0, 5)
	fmt.Println(array)
}

func QuickSort(array []int, L int, R int) {
	if L >= R {
		return
	}

	left := L
	right := R

	// ⚠️ 要取 array[left]，不要取 array[0]
	pivot := array[left]

	for left < right {
		// 第一步，把比 pivot 大的值放到 pivot 右边
		// 如果右指针的值比选定值大，右指针左移
		for left < right && array[right] >= pivot {
			right --
		}
		// 上面的循环退出来，说明右侧的值比 pivot 要小，将小的值换到左指针位置
		if left < right {
			array[left] = array[right]
		}

		// 第二步，把比 pivot 小的值放到 pivot 左边
		// 如果左指针的值比选定值小，左指针右移
		for left < right && array[left] <= pivot {
			left ++
		}
		// 上面的循环退出来，说明左侧的值比 pivot 要大，将大的值换到右指针位置
		if left < right {
			array[right] = array[left]
		}

		// 将 pivot 放到合适的位置
		if left >= right {
			array[left] = pivot
		}
	}

	QuickSort(array, L, right-1)
	QuickSort(array, right+1, R)
}
