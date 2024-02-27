package array

import "testing"

func TestSearch(t *testing.T) () {
	t_A := []int{1, 2, 3, 4, 5}
	target := -2

	result := Search(t_A, target)
	t.Logf("%d", result)
}
