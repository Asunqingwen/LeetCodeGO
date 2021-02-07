package Array

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 || index >= len(nums) {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}

	c = append(c, nums[index])
	findcombinationSum(nums, target-nums[index], index, c, res)
	c = c[:len(c)-1]
	findcombinationSum(nums, target, index+1, c, res)
}
