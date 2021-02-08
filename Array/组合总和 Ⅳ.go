package Array

func combinationSum4(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for j := 0; j < len(nums); j++ {
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}
