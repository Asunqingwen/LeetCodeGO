package String

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	len_ := len(s)
	if len_ < 2 {
		return len_
	}
	var freq [256]int
	res, left, right := 1, 0, 1
	freq[s[0]-'a'] = 1
	for right < len_ {
		if freq[s[right]-'a'] == 0 {
			freq[s[right]-'a']--
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		res = max(res, right-left)
	}
	return res
}
