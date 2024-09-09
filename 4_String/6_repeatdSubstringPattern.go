package string

// leetcode:459
func repeatedSubstringPattern(s string) bool {
	result := false
	for i := 1; i <= len(s)/2; i++ {
		for j := 0; j <= len(s)-i; j+=i {
			if s[j:j+i] != s[:i] {
				break
			}
			if j == len(s)-i {
				result = true
			}
		}
	}
	return result
}