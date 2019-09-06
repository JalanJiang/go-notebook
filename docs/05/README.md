## 5.0 控制结构

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