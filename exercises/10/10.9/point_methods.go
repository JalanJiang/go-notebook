/*
从 point.go 开始（第 10.1 节的练习）：使用方法来实现 Abs() 和 Scale()函数，Point 作为方法的接收者类型。也为 Point3 和 Polar 实现 Abs() 方法。完成了 point.go 中同样的事情，只是这次通过方法。
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

type Point3 struct {
	X, Y, Z float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

type Polar struct {
	R, T float64
}

func (p Polar) Abs() float64 {
	return p.R
}

func main() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs())

	p2 := &Point{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", p2.Abs())

	p1.Scale(5)
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", p1.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f", p1.X, p1.Y)
}

/*Output:
The length of the vector p1 is: 5.000000
The length of the vector p2 is: 6.403124
The length of the vector p1 after scaling is: 25.000000
Point p1 after scaling has the following coordinates: X 15.000000 - Y 20.000000
*/
