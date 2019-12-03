/*
定义结构体 employee，它有一个 salary 字段，
给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水。
*/
package main

import "fmt"

type employee struct {
	salary float32
}

// 按照百分比增加薪水
func (this *employee) giveRaise(pct float32) {
	this.salary += this.salary * pct
}

func main() {
	var e = new(employee)
	e.salary = 10000
	e.giveRaise(0.4)
	fmt.Println(e.salary)
}

/*Output:
14000
*/
