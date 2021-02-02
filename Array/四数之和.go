package Array

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result, first, second, third, four, length := make([][]int, 0), 0, 0, 0, 0, len(nums)

	for first = 0; first < length-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second = first + 1; second < length-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			four = length - 1
			ans := target - nums[first] - nums[second]
			for third = second + 1; third < length-1; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				for third < four && nums[third]+nums[four] > ans {
					four--
				}
				if third == four {
					break
				}
				if nums[third]+nums[four] == ans {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
		}
	}
	return result
}
