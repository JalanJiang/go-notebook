// 使用 * 符号打印宽为 20，高为 10 的矩形。
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
