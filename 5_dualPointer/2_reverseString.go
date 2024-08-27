package dualPointer

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		temp := s[right]
		s[right] = s[left]
		s[left] = temp

		right --
		left ++
	}
}