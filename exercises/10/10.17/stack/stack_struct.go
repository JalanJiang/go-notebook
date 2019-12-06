/*
使用包含一个索引和一个 int 数组的结构体作为底层数据结构，索引表示第一个空闲的位置。
*/
package stack

import (
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
