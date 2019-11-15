// main 函数中写一个用于打印 Hello World 字符串的匿名函数并赋值给变量 fv，然后调用该函数并打印变量 fv 的类型。
package main

import (
	"fmt"
)

func main() {
	fv := func() {
		fmt.Print("Hello World!")
	}
	fv()
	fmt.Printf("The type of fv is %T", fv)
}

// Hello World!The type of fv is func()
