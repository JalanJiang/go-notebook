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