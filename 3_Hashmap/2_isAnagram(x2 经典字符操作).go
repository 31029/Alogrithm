package hashmap

func IsAnagram(s string, t string) bool {
	// x1: 未考虑两个字符串不同长度的情况
	if len(s) != len(t) {
		return false
	}
	s_map := make(map[rune]int)
	t_map := make(map[rune]int)

	for _, v := range s {
		s_map[v]++
	}
	for _, v := range t {
		t_map[v]++
	}

	for k := range s_map {
		if s_map[k] != t_map[k] {
			return false
		}
	}
	return true
}



func isAnagram_classic(s string, t string) bool {
    record := [26]int{}
    
    for _, r := range s {
        record[r-rune('a')]++
    }
    for _, r := range t {
        record[r-rune('a')]--
    }

    return record == [26]int{}   // record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
}