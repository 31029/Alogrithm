package backtracking

import (
	"strconv"
)

// leetcode-93: 复原 IP 地址 [https://leetcode.cn/problems/restore-ip-addresses/description/]
/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 part ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 part 中插入 '.' 来形成。
你 不能 重新排序或删除 part 中的任何数字。你可以按 任何 顺序返回答案。

示例：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
*/
var result = []string{}
var curpath = []string{}


func isLegalPart(part string) bool {
	if len(part) > 1 && part[0] == '0' {
		return false
	}

	ipart, err := strconv.Atoi(part)
	if err == nil {
		if ipart >= 0 && ipart <= 255 {
			return true
		}
	}

	return false
}

func restoreIpAddressesBackTracking(part string, raw string, length, height int) {
    if height == 4 {
		if isLegalPart(part) && length == len(raw) {
			path := ""
			for i, part := range curpath {
				if i != 3 {
					path += part + "."
				} else {
					path += part
				}
			}
			result = append(result, path)
		}
		return
	}

	for i := 0; i < 3; i++ {
		if length + i + 1 > len(raw) {
			return
		}
		newpart := raw[length: length+i+1]
		newlen := length+len(newpart)
		
		// 剪枝, 注意此处其实是 4 - (height+1)
		if len(raw) - newlen > (4 - height - 1) * 3 {
			continue
		}
		if isLegalPart(newpart) {
			curpath = append(curpath, newpart)
			restoreIpAddressesBackTracking(newpart, raw, newlen, height+1)
			curpath = curpath[:len(curpath)-1]
		}
	}
}


func RestoreIpAddresses(s string) []string {
	// 思路： 经典回溯方法解决，注意剪枝. 
	// 技巧：ip地址由四个段组成，每个段的长度为 0~3，由此构建回溯中的for循环

	// 每次调用都要初始化 result 数组
	result = []string{}
	// 注意初始化条件， height = 0
	restoreIpAddressesBackTracking("", s, 0, 0)
	return result
}


