package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8) // 返回介于 [0, n) 之间的伪随机数
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	times := int64(time.Now().Nanosecond())
	rand.Seed(times) // 使用 Seed(value) 函数来提供伪随机数的生成种子
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32()) // 返回介于 [0.0, 1.0) 之间的伪随机数
	}
}
