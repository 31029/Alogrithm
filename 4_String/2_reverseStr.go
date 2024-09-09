package string

// leetcode: 541. 反转字符串 II
func ReverseStr(s string, k int) string {
	count := len(s) / (2*k)
	left := len(s) % (2*k)

	byte_s := []byte(s)

	for i := 0; i < count; i++ {
		ReverseString(byte_s[i*2*k: i*2*k + k])
	}
	if left >= k {
		ReverseString(byte_s[count*2*k: count*2*k + k])
	} else {
		ReverseString(byte_s[count*2*k: count*2*k + left])
	}

	return string(byte_s)

}