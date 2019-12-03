package main

import (
	"fmt"
)

type TestStruct struct {
	a float64
	int
	string
}

func main() {
	// 字面量创建
	test := TestStruct{3.0, 3, "test"}
	fmt.Println(test.a, test.int, test.string)
	fmt.Println(test)

	p := Person{"a", "b"}
	fmt.Println(p.lastName)
}

/*Output:
3 3 test
{3 3 test}
*/
