package string

import "testing"

func TestReplaceNumber(t *testing.T) {
	s := "a7r5"
	s1 := ReplaceNumber(s)

	t.Logf("%v", s1)
}