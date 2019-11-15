// 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序。
package main

import (
	"fmt"
)

// 闭包：返回值是函数
func fib() func() int {
	a, b := 1, 1
	// 无论外部函数退出与否，闭包内总能操作外部函数局部变量，并保留值
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()

	for i := 0; i <= 10; i++ {
		fmt.Println(i+2, f())
	}
}

/* Output:
2 2
3 3
4 5
5 8
6 13
7 21
8 34
9 55
10 89
11 144
12 233
*/
