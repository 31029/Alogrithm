package dualPointer

// " the sky is blue  "
func ReverseWords(s string) string {
	// 快慢指针：移除前面、中间、后面存在的多余空格 !!!
	// 之后的思路：整个反转字符串之后再逐一反转其中单词。
	// 个人觉得hashmap的思路更加清晰。
	b := []byte(s)
    slow := 0
    for i := 0; i < len(b); i++ {
        if b[i] != ' ' {
			// 加入中间的空格
            if slow != 0 {
                b[slow] = ' '
                slow++
            }
            for i < len(b) && b[i] != ' ' { // 复制逻辑
                b[slow] = b[i]
                slow++
                i++
            }
        }
    }

	// 最后剪裁
    b = b[0:slow]

	// 先整个反转
	ReverseString(b)
	s = string(b)

	// 再逐一反转其中每一个单词
	slow, fast := 0, 0
	for fast < len(b) {
		if s[fast] == ' ' {	
			ReverseString(b[slow: fast])
			fast ++
			slow = fast
			continue
		}
		fast ++ 
	}

	// 反转最后一个单词
	ReverseString(b[slow: fast])
	return string(b)
}