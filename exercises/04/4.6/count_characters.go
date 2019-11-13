package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	string1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string1 is %d\n", len(string1))
	// RuneCountInString 获得字符数量
	fmt.Printf("The number of characters in string1 is %d\n", utf8.RuneCountInString(string1))

	string2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string2 is %d\n", len(string2))
	fmt.Printf("The number of characters in string2 is %d\n", utf8.RuneCountInString(string2))
}

/* Output:
The number of bytes in string1 is 22
The number of characters in string1 is 22
The number of bytes in string2 is 28
The number of characters in string2 is 24
*/
