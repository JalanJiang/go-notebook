// 以下程序输出什么？
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		var v int // 重新声明了，没有赋值 = 0
		fmt.Printf("%d ", v)
		v = 5
	}
}

/* Output:
0 0 0 0 0
*/
