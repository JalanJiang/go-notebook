// 写一个 Season 函数，要求接受一个代表月份的数字，然后返回所代表月份所在季节的名称（不用考虑月份的日期）。
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Season(11))
}

// 实现一个 season 函数
func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "冬天"
	case 3, 4, 5:
		return "春天"
	case 6, 7, 8:
		return "夏天"
	case 9, 10, 11:
		return "秋天"
	}
	return "未知季节"
}
