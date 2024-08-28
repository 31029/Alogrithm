package dualPointer

// " the sky is blue  "
func ReverseWords(s string) string {
	// 快慢指针：移除前面、中间、后面存在的多余空格 !!!
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

	sb := b 
	ReverseString(sb)
	s = string(sb)

	slow, fast := 0, 0
	for fast < len(sb) {
		if s[fast] == ' ' {	
			ReverseString(sb[slow: fast])
			fast ++
			slow = fast
			continue
		}
		fast ++ 
	}
	ReverseString(sb[slow: fast])
	return string(sb)
}