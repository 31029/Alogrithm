package string

func reverseWords(s string) string {
	h_map := make(map[int]string)
	temp := ""
	wordCount := 0
	for _, v := range s {
		if v == ' ' {
			if temp != "" {
				h_map[wordCount] = temp
				temp = ""
				wordCount ++
			}
			continue
		}
		temp += string(v)
	}
	// 如果结尾没有空格，加入最后一个单词
	if temp != "" {
		h_map[wordCount] = temp
		wordCount ++
	}

	result := ""
	for i := wordCount-1; i >= 0; i-- {
		if i == 0 {
			result = result + h_map[i]
			return result
		}
		result = result + h_map[i] + " "
	}
	return result
}