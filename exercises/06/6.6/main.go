/*
实现一个输出前 30 个整数的阶乘的程序。

n! 的阶乘定义为：n! = n * (n-1)!, 0! = 1，因此它非常适合使用递归函数来实现。

然后，使用命名返回值来实现这个程序的第二个版本。

特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误。那么，如何才能解决这个问题呢？

最好的解决方案就是使用 big 包（详见第 9.4 节）。
*/

package main

import "fmt"

func main() {
	fmt.Print(Factorial(30))
	fmt.Print(Factorial2(30))
}

// 使用 uint64
func Factorial(n uint64) (fac uint64) {
	if n > 0 {
		fac = n * Factorial(n-1)
	} else {
		fac = 1
	}
	return
}

func Factorial2(n uint64) uint64 {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}
