package main

import (
	"fmt"

	"./stack"
)

func main() {
	st := new(stack.Stack)
	st.Push(3)
	fmt.Printf("%v\n", st)
}
