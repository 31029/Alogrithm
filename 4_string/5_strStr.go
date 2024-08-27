package string

func hasPrefix(src, tgt string) bool {
	return len(src) >= len(tgt) && src[:len(tgt)] == tgt
}


// leetcode:28 
func strStr(haystack string, needle string) int {
	for i := range haystack {
		if hasPrefix(haystack[i:], needle) {
			return i
		}
	}
	return -1
}