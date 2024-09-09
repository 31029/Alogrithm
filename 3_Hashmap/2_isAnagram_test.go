package hashmap

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s := "abcd"
	g := "dcc"

	t.Logf("%t", IsAnagram(s, g))
}
