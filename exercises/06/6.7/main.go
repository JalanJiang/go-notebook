/*
包 strings 中的 Map 函数和 strings.IndexFunc() 一样都是非常好的使用例子。
请学习它的源代码并基于该函数书写一个程序，要求将指定文本内的所有非 ASCII 字符替换成 ? 或空格。
您需要怎么做才能删除这些字符呢？
*/
// strings_map.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// rune 类型
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	// 对每一个字符执行 asciiOnly 函数
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
}
