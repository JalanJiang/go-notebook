/*
定义一个 Rectangle 结构体，它的长和宽是 int 类型，并定义方法 Area() 和 Perimeter()，然后进行测试。
*/
package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", r1.Area())
	fmt.Println("Rectangle perimeter is: ", r1.Perimeter())
}

/* Output:
Rectangle is:  {4 3}
Rectangle area is:  12
Rectangle perimeter is:  14
*/
