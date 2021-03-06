// 使用按位补码从 0 到 10，使用位表达式 %b 来格式化输出。
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%b %b\n", i, ^i)
	}
}

/* Output
0 -1
1 -10
10 -11
11 -100
100 -101
101 -110
110 -111
111 -1000
1000 -1001
1001 -1010
*/
