package hashmap

func CanConstruct(ransomNote string, magazine string) bool {
	n_map := make(map[rune]int, len(ransomNote))
	for _, v := range ransomNote {
		n_map[v] ++
	}

	for _, v := range magazine {
		n_map[v]  --
	}

	for _, v := range n_map {
		if v > 0 {
			return false			
		}
	}

	return true
}