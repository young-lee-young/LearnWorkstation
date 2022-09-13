package main

/**
	多数元素，数量大于一半的数 LeetCode No169

	解题思路：不相同就同时删除，相同长度加1，最后留下的数就是多数元素
 */
func main() {
	numSlice := []int{2,2,1,1,1,2,2}
	num := numSlice[0]
	count := 1
	for i := 1; i < len(numSlice); i ++ {
		if numSlice[i] == num {
			count ++
		} else {
			if count > 0 {
				count --
			} else {
				num = numSlice[i]
				count = 1
			}
		}
	}
}
