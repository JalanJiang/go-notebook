/*
使用包含一个索引和一个 int 数组的结构体作为底层数据结构，索引表示第一个空闲的位置。
*/

package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix == LIMIT {
		return // 满了
	}
	st.data[st.ix] = n
	st.ix++ // 支持++
}

func (st *Stack) Pop() int {
	if st.ix != 0 {
		st.ix--
		return st.data[st.ix]
	}
	return 0
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "] "
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
