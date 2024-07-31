package string

func ReverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		t := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = t
	}
}
