/*
为 int 定义一个别名类型 Day，定义一个字符串数组它包含一周七天的名字，为类型 Day 定义 String() 方法，它输出星期几的名字。
使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）。
*/
package main

import "fmt"

type Day int

// 枚举
const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

// 数组，对应名称
var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func main() {
	var th Day = 3
	fmt.Printf("The 3rd day is: %s\n", th)

	var day = SU
	fmt.Println(day)
	fmt.Println(0, MO, 1, TU)
}

/*
The 3rd day is: Thursday
Sunday
0 Monday 1 Tuesday
*/
