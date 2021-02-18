package String

func generateParenthesis(n int) []string {
	s, res := "", []string{}
	generateParenthese(n, 0, 0, s, &res)
	return res
}

func generateParenthese(n, lc, rc int, s string, res *[]string) {
	if len(s) == 2*n {
		*res = append(*res, s)
		return
	}
	if lc < n {
		s += "("
		generateParenthese(n, lc+1, rc, s, res)
		s = s[:len(s)-1]
	}
	if rc < lc {
		s += ")"
		generateParenthese(n, lc, rc+1, s, res)
		s = s[:len(s)-1]
	}
}
