package main

import "fmt"

/**
	无序数组两数之和返回下标 LeetCode No001

	解题思路：利用map，键为值、值索引

	标签：哈希表

	时间复杂度：O(n)

	空间复杂度：O(n)
 */
func main() {
	numArray := []int{2, 7, 11, 15}
	sum := 9

	hashMap := make(map[int]int)
	for i := 0; i < len(numArray); i ++ {
		num := sum - numArray[i]
		if value, ok := hashMap[num]; ok {
			fmt.Println(value, i)
			break
		}
		hashMap[numArray[i]] = i
	}
}

/**
	第二种方法：先排序 + 双指针

	时间复杂度：O(nlogn)

	空间复杂度：O(1)
 */

/**
问题1：如果输入数据已经排序，哪个算法（哈希表 VS 双指针）更好？

双指针更好，时间复杂度为 O(n)，空间复杂度为 O(1)


问题2：如果需要返回所找的两个数在数组中的下标，哪个算法（哈希表 VS 双指针）更新？

哈希表更好
 */
