package hashmap

import "strconv"


// 经典错误: 还没想清楚思路就开始动笔。↓ ↓ ↓ ↓ ↓

// func getNewN(n int) int {
//	// 用数学方法读取每位数字
// 	tempMap := make(map[int]int)
// 	cur_i := -1
// 	cur_index := 10

// 	// 先计算一次
// 	cur_i = n / cur_index

// 	// 计算非个位的取值
// 	for cur_i != 0 {
// 		tempMap[cur_index] = cur_i
// 		cur_index = cur_index * 10
// 		cur_i = n / cur_index
// 	}

// 	// 计算个位的值
// 	n1 := n
// 	for k, v := range tempMap {
// 		n1 = n1 - k*v
// 	}
// 	tempMap[1] = n1

// 	newN := 0
// 	for _, v := range tempMap {
// 		newN += v * v
// 	}
// 	return newN
// }


func getNewN(n int) int {
	sum := 0
	// 将整数转换为字符串
	strNum := strconv.Itoa(n)
	for _, digit := range strNum {
		// 将字符转换为整数
		num, _ := strconv.Atoi(string(digit))
		// 计算平方并累加到sum
		sum += num * num
	}
	return sum
}

func IsHappy(n int) bool {
	iMap := make(map[int]int)
	iMap[n] ++

	n = getNewN(n)
	for n != 1 {
		n = getNewN(n)
		if iMap[n] == 0 {
			iMap[n] ++
		} else {
			return false
		}
	}

	return true
}