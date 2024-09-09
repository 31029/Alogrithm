package string

import "testing"

func TestReverseString(t *testing.T) {

	// s := []byte("hello")

	s_str := "world"
	s := []byte(s_str)
	ReverseString(s)

	t.Logf("result: %s", string(s))

}
