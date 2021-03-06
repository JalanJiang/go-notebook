## 6.1 介绍

- Go是编译型语言，所以函数编写的顺序是无关紧要的
- 在 Go 里面函数重载是不被允许的
  - 函数重载（function overloading）指的是可以编写多个同名函数
- Go 语言不支持这项特性的主要原因是函数重载需要进行多余的类型匹配影响性能；没有重载意味着只是一个简单的函数调度
  
如果需要申明一个在外部定义的函数，你只需要给出函数名与函数签名，不需要给出函数体：

```go
func flushICache(begin, end uintptr) // implemented externally
```

函数也可以以申明的方式被使用，作为一个函数类型，就像：

```go
type binOp func(int, int) int
```

----

## 6.2 函数参数与返回值

- 任何一个有返回值（单个或多个）的函数都必须以 return 或 panic（参考 第 13 章）结尾
- 函数定义时，它的形参一般是有名字的，不过我们也可以定义没有形参名的函数 `func f(int, int, float64)`

### 6.2.1 按值传递（call by value） 按引用传递（call by reference）

- Go 默认使用按值传递来传递参数，也就是传递参数的副本
- 如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址 `Function(&arg1)`
- 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）

### 6.2.2 命名的返回值（named return variables）

即使只有一个命名返回值，也需要使用 `()` 括起来。

```go
// 命名的返回值
func get(input int) (x1 int) {
    // ....
    x1 = 1
    return
}
```

> 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。

### 6.2.3 空白符（blank identifier）

空白符 `_` 用来匹配一些不需要的值，然后丢弃掉。

### 6.2.4 改变外部变量（outside variable）

传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 `return` 返回。

```go
// reply 是一个 int 类型指针
func Multiply(a, b int, reply *int) {
    // *reply 取 reply 地址指向的值
    *reply = a * b
}

func main() {
    n := 0
    // 拿到 n 的地址
    reply := &n
    Multiply(10, 5, reply)
}
```

----

## 6.3 传递变长参数

如果函数的最后一个参数是采用 `...type` 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变参函数。

```go
func Greeting(prefix string, who ...string)
// 变量 who 的值为 []string{"Joe", "Anna", "Eileen"}
Greeting("hello", "Joe", "Anna", "Eileen")
```

如果参数被存储在一个 `slice` 类型的变量 `slice` 中，则可以通过 `slice...` 的形式来传递参数，调用变参函数：

```go
// ???
func main() {
    slice := []int{1, 2, 3}
    x = min(sloce...)
}

func min(s ...int) int {
    // ...
}
```

**但是如果变长参数的类型并不是都相同的：**

方案一：使用结构

```go
type Options struct {
    par1 type1,
    par2, type2,
    // ...
}
```

方案二：使用空接口

如果一个变长参数的类型没有被指定，则可以使用默认的空接口 `interface{}`，这样就可以接受任何类型的参数（详见第 11.9 节）。

----

## 6.4 defer 和追踪

- 关键字 `defer` 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
- 一般用于释放某些已分配的资源
- 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）

关键字 defer 允许我们进行一些函数执行完成后的收尾工作：

1. 关闭文件流
2. 解锁一个加锁的资源
3. 打印最终报告
4. 关闭数据库连接
5. 实现代码追踪：一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息

----

## 6.5 内置函数

- close
- len
- cap
- new
- make
- copy
- append
- panic
- recover
- print
- println
- complex
- real imag

----

## 6.6 递归函数

当一个函数在其函数体内调用自身，则称之为递归。

在使用递归函数时经常会遇到的一个重要问题就是栈溢出：一般出现在大量的递归调用导致的程序栈内存分配耗尽。这个问题可以通过一个名为 [懒惰求值](https://zh.wikipedia.org/wiki/%E6%83%B0%E6%80%A7%E6%B1%82%E5%80%BC) 的技术解决。

----

## 6.7 将函数作为参数

函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为**回调**。

```go
func main() {
    callback(1, Add)
}

func Add(a, b int) {
    // do something
}

func callback(y int, f func(int, int)) {
    f(y, 2)
}
```

----

## 6.8 闭包

- 当我们不希望给函数起名字的时候，可以使用匿名函数
- 匿名函数不能够独立存在，但可以被赋值于某个变量，即保存函数的地址到变量中，然后通过变量名对函数进行调用
- 匿名函数同样被称之为闭包
  - 被允许调用定义在其它环境下的变量，可使得某个函数捕捉到一些外部状态
  - 一个闭包继承了函数所声明时的作用域
  - 这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁

----

## 6.9 应用闭包：将函数作为返回值

```go
package main

import "fmt"

func main() {
	// p2 是一个 +2 的函数
	p2 := Add2()
	fmt.Printf("call Add2 for 3 gives:%v\n", p2(3))

	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

// Add2 is a function
func Add2() func(b int) int {
	// 返回一个函数
	return func(b int) int {
		return b + 2
	}
}

// Adder is a function
func Adder(a int) func(b int) int {
	// 返回一个函数
	return func(b int) int {
		return a + b
	}
}

// Output:
// call Add2 for 3 gives:5
// The result is: 5
```

闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够**继续操作外部函数中的局部变量**。在多次调用中，变量 x 的值是被保留的：

```go
package main

import "fmt"

func main() {
	// p2 是一个 +2 的函数
	var f = Adder()
	fmt.Print(f(1), "-")
	fmt.Print(f(20), "-")
	fmt.Print(f(300))
}

// Adder is a function
func Adder() func(b int) int {
	// 返回一个函数
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// Output:
// 1 - 21 - 321
```

----

## 6.10 使用闭包调试

可以使用 `runtime` 或 `log` 包中的特殊函数来实现调试功能。

```go
package main

import (
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
}

// 2019/09/06 18:30:45 /Users/jalan/www/own/go-demo/main.go:18
```

----

## 6.11 计算函数执行时间

```go
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

----

## 6.12 通过内存缓存来提升性能

当在进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。（其实就是存在数组/map……里）