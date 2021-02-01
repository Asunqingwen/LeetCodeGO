package main

func removeElement(nums []int, val int) int {
	first, second := 0, 0
	for second < len(nums) {
		if nums[first] == val {
			nums[first], nums[second] = nums[second], nums[first]
		}
		if nums[first] != val {
			first++
		}
		second++
	}
	return first
}
