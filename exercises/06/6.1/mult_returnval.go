// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
package main

import (
	"fmt"
)

func main() {
	e, f, g := Operate(2, 3)
	fmt.Printf("%d %d %d\n", e, f, g)
	e, f, g = Operate2(3, 4) // 注意第二次赋值
	fmt.Printf("%d %d %d\n", e, f, g)
}

func Operate(a int, b int) (int, int, int) {
	return a + b, a * b, a / b
}

func Operate2(a int, b int) (c int, d int, e int) {
	c, d, e = a+b, a*b, a/b
	return
}
