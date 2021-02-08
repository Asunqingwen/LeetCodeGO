package Array

func combinationSum3(k, n int) [][]int {
	c, res := []int{}, [][]int{}
	findcombinationSum3(0, k, n, c, &res)
	return res
}

func findcombinationSum3(index, k, n int, c []int, res *[][]int) {
	if k < len(c) {
		return
	}
	if n == 0 && k == len(c) {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}

	for i := index; i < 9; i++ {
		if n < i+1 {
			break
		}
		c = append(c, i+1)
		findcombinationSum3(i+1, k, n-i-1, c, res)
		c = c[:len(c)-1]
	}
	return
}
