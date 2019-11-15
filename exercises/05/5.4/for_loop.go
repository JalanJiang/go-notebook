// 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值。
package main

import (
	"fmt"
)

func main() {
	// 计数器
	var count int
	count = 0

	for i := 0; i < 15; i++ {
		count++
	}

	fmt.Printf("计数器的值 = %d\n", count)

	// 使用 goto 语句重写循环，要求不能使用 for 关键字
	count = 0

LOOP:
	count++
	if count == 15 {
		fmt.Printf("计数器的值 = %d\n", count)
		return
	}
	goto LOOP
}
