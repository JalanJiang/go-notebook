## 10.1 结构体定义

### 定义

- 如果字段在代码中从来也不会被用到，那么可以命名它为 `_`。

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```

### 声明

```go
var s T
s.a = 5
s.b = 8
```

### 使用 new

使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：`var t *T = new(T)`。

```go
var t *T
t = new(T)
```

### 初始化

```go
ms := &struct1{10, 15.5, "Chris"}
// 或
var ms struct1
ms = struct1{10, 15.5, "Chris"}
```

几种初始化方式：

```go
type Interval struct {
    start int
    end   int
}

intr := Interval{0, 3}            
intr := Interval{end:5, start:1}  
intr := Interval{end:5}   
```

### 结构体的内存布局

- Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

![](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/10.1_fig10.2.jpg?raw=true)

### 递归结构体

结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素（通常叫节点）时特别有用，此时节点包含指向临近节点的链接（地址）。

1. 链表的 `next` 节点指向下一个 `Node` 结构体
2. 二叉树的 `left` / `right` 节点指向下一个 `Node` 结构体

### 结构体转换

Go 中的类型转换遵循严格的规则。当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型，它们可以互相转换。

```go
package main
import "fmt"

type number struct {
    f float32
}

type nr number   // alias type

func main() {
    a := number{5.0}
    b := nr{5.0}
    // var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
    // var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
    // var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
    // needs a conversion:
    var c = number(b)
    fmt.Println(a, b, c)
}
```