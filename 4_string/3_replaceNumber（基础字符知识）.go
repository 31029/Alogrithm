package string

func ReplaceNumber(s string) string {
	head := ""
	for _, str := range s {
		if str <= '9' && str >= '0'  {
			head += "number"
		} else {
			head += string(str)
		}
	}
	return head
}