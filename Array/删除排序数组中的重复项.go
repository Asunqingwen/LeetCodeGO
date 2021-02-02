package Array

func removeDuplicates(nums []int) int {
	first, second := 0, 0
	for second < len(nums) {
		if nums[first] != nums[second] {
			first++
			nums[first] = nums[second]
		}
		second++
	}
	return first + 1
}
