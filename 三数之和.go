package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, length := make([][]int, 0), 0, 0, 0, len(nums)
	if length > 0 && (nums[0] > 0 || nums[length-1] < 0) {
		return result
	}

	for index = 0; index < length; index++ {
		start, end = 0, length-1
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		end = length - 1
		target := -nums[index]
		for start = index + 1; start < length; start++ {
			if start > index+1 && nums[start] == nums[start-1] {
				continue
			}
			for start < end && nums[start]+nums[end] > target {
				end--
			}
			if start == end {
				break
			}
			if nums[start]+nums[end] == target {
				result = append(result, []int{nums[index], nums[start], nums[end]})
			}
		}
	}
	return result
}
