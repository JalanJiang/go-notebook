## 5.1 if-else 结构

当 if 结构内有 break、continue、goto 或者 return 语句时，Go 代码的常见写法是省略 else 部分。

```go
if condition {
    return x
}
return y
```

if 可以包含一个初始化语句（如：给一个变量赋值）：

```go
if initialization; condition {
    // do something
}
```

例如：

```go
if val := 10, val > max {
    // do something
}
```

----

## 5.2 测试多返回值函数的错误

Go 语言的函数经常使用两个返回值来表示执行是否成功：

- 返回某个值以及 true 表示成功
- 返回零值（或 nil）和 false 表示失败
- 当不使用 true 或 false 的时候，也可以使用一个 error 类型的变量来代替作为第二个返回值：成功执行的话，error 的值为 nil，否则就会包含相应的错误信息

```go
if value, ok := readData(); ok {
    // do something
}
```

----

## 5.3 switch 结构

- `var1` 可以是任何类型
- `val1` 和 `val2` 必须是相同类型
- 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 `switch` 代码块，不需要特别使用 `break` 语句来表示结束

```go
switch var1 {
    case val1:
        // ...
    case val2:
        // ...
    default:
        // ...
}
```

希望继续执行后续分支的代码，可以使用 `fallthrough` 关键字来达到目的：

```go
switch i {
    // 当代码块只有一行时，可以直接放置在 case 语句之后
    case 0: fallthrough
    case 1:
        f()
}
```