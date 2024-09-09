package string

import "testing"

func TestReverseStr(t *testing.T) {

	// s := []byte("hello")

	s_str := "abcdef"

	s := ReverseStr(s_str, 3)

	t.Logf("result: %s", s)

}