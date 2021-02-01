package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] == nums[left] {
				left++
			}
			if nums[mid] == nums[right] {
				right--
			}
		}
	}
	return -1
}
