// 使用递归函数从 10 打印到 1。
package main

import "fmt"

func main() {
	PrintNum(10)
}

func PrintNum(num int) {
	if num < 1 {
		return
	}
	fmt.Println(num)
	PrintNum(num - 1)
}
