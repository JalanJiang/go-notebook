/*
定义一个结构体类型 Base，它包含一个字段 id，方法 Id() 返回 id，方法 SetId() 修改 id。
结构体类型 Person 包含 Base，及 FirstName 和 LastName 字段。
结构体类型 Employee 包含一个 Person 和 salary 字段。

创建一个 employee 实例，然后显示它的 id。
*/
package main

import "fmt"

type Base struct {
	id int64
}

func (b *Base) Id() int64 {
	return b.id
}

func (b *Base) SetId(id int64) {
	b.id = id
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float64
}

func main() {
	employee := new(Employee)
	employee.SetId(1)
	employee.FirstName = "Jalan"
	employee.LastName = "Jiang"
	employee.salary = 200000.0

	fmt.Println(employee)

	// 嵌套
	idjb := Base{9}
	jb := Person{idjb, "James", "Bond"}
	e := &Employee{jb, 10000.}
	fmt.Printf("ID of our hero: %v\n", e.Id())
	e.SetId(10)
	fmt.Printf("The new ID of our hero: %v\n", e.Id())
}

/*Output:
&{{{1} Jalan Jiang} 200000}
ID of our hero: 9
The new ID of our hero: 10
*/
