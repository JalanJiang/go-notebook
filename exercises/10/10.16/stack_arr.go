package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack [LIMIT]int

func (st *Stack) Push(n int) {
	for ix, v := range st {
		// 位置为 0 表示可以 push 数
		if v == 0 {
			st[ix] = n
		}
	}
}

func (st *Stack) Pop() int {
	end := len(st) - 1
	// 弹出的数置 0
	for ix := end; ix >= 0; ix-- {
		if v := st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}

	return 0
}

func (st Stack) String() string {
	str := ""
	for ix, v := range st {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "] "
	}
	return str
}

func main() {
	st1 := new(Stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}

/*
[0:3] [1:3] [2:3] [3:3]
[0:3] [1:3] [2:3] [3:3]
[0:3] [1:3] [2:3] [3:3]
Popped 3
[0:3] [1:3] [2:3] [3:0]
Popped 3
[0:3] [1:3] [2:0] [3:0]
Popped 3
[0:3] [1:0] [2:0] [3:0]
Popped 3
[0:0] [1:0] [2:0] [3:0]
*/
