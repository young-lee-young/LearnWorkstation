package binary

/**
二分搜索
*/
func binarySearch(nums []int, target int) int {
	// 在[start, end]范围里查找
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start+1)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
