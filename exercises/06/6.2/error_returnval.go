// 编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，
// 如果参数是一个负数的话将返回一个错误。
// 编写两个版本，一个是非命名返回值，一个是命名返回值。
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Print(MySqrt(4))
}

func MySqrt(num float64) (float64, error) {
	if num < 0 {
		return float64(math.NaN()), errors.New("错误信息")
	}
	return math.Sqrt(num), nil
}

func MySqrtN(num float64) (rNum float64, err error) {
	if num < 0 {
		rNum = float64(math.NaN())
		err = errors.New("错误信息")
	} else {
		rNum = math.Sqrt(num)
		// err = nil // 可以不用赋值
	}
	return
}
