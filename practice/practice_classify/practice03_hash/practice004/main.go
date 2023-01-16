package main

/**
LeetCode No001 无序数组两数之和返回下标

解题思路：利用map，键为值、值索引

标签：哈希表

时间复杂度：O(n)

空间复杂度：O(n)
*/
func main() {
	numArray := []int{2, 7, 11, 15}
	sum := 9

	hashMap := make(map[int]int)
	for i := 0; i < len(numArray); i++ {
		num := sum - numArray[i]
		if _, ok := hashMap[num]; ok {
			break
		}
		hashMap[numArray[i]] = i
	}
}
