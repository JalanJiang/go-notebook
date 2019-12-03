package main

import (
	"fmt"
	"reflect"
)

func main() {

	testString := "Hello，世界"
	// fmt.Println(string([]rune(testString)[:2])) // 输出：「你好」

	// for i := 0; i < len(testString); i++ {
	// 	c := testString[i]
	// 	fmt.Printf("%c 的类型是 %s\n", c, reflect.TypeOf(c))
	// }

	for _, c := range testString {
		fmt.Printf("%c 的类型是 %s\n", c, reflect.TypeOf(c))
	}

	fmt.Printf("字符 %c 对应的整型为 %d\n", byteC, byteC)
	// Output: 字符 j 对应的整型为 106
}
