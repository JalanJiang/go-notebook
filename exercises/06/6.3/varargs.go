// 写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。

package main

import (
	"fmt"
)

func main() {
	PrintElement("a", "b", "c")
}

func PrintElement(s ...string) {
	for _, e := range s {
		fmt.Println(e)
	}
}
