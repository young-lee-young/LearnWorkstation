/**
 * @Time:    2021/3/31 22:48 
 * @Author:  leeyoung
 * @File:    reverse.go
 * @Content: 
 */
package reverse

// 反转切片
func ReverseSlice(nums []int) []int {
	for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
