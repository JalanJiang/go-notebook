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

----

## 5.4 for 结构

计数器迭代：

```
for 初始化语句; 条件语句; 修饰语句 {}
```

使用多个计数器：

```go
for i, j := 0, N; i < j; i, j = i+1, j-1 {}
```

### 5.4.2 基于条件判断的迭代

没有头部的条件判断迭代（类似其它语言中的 while 循环）：

```
for 条件语句 {}
```

```go
for i >= 0 {
    // dosomething
}
```

### 5.4.3 无限循环

如果 for 循环的头部没有条件语句，那么就会认为条件永远为 true：

```go
for {
    // 无限循环
}
```

### 5.4.4 for-range 结构

- 可以迭代任何一个集合
- 语法上很类似其它语言中 `foreach` 语句

```
for ix, val := range coll { }
```

- `val` 始终为集合中对应索引的值拷贝，一般只具有只读性质

迭代字符串：

```go
for pos, char := range str {
    // ...
}
```

----

## 5.5 Break 与 continue

- break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环
  - 在 switch 或 select 语句中（详见第 13 章），break 语句的作用结果是跳过整个代码块，执行后续的代码
- continue 忽略剩余的循环体而直接进入下一次循环的过程