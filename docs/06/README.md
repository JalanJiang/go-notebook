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