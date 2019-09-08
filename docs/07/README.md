## 7.1 声明和初始化

### 7.1.1 概念

- 数组是具有相同**唯一类型**的一组**已编号**且**长度固定**的数据项序列
- 如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型

```go
// 声明方式
// var identifier [len]type
var arr1 [5]int // 当声明数组时所有的元素都会被自动初始化为默认值 0
```

for-range 遍历：

```go
// i 是数组索引
for i, _ := range arr1 {
    // do something
}
```

Go 语言中的数组是一种**值类型**（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建：

```go
var arr2 = new([5]int)
```

`arr2` 与 `arr1` 的区别：

- `arr1` 的类型是 `*[5]int`
- `arr2` 的类型是 `[5]int`

这样的结果就是当把一个数组赋值给另一个时，需要在做一次数组内存的拷贝操作：

```go
arr1 := *arr2
arr1[2] = 100  // 在赋值后修改 arr1 不会对 arr2 生效
```

### 7.1.2 数组常量

```go
// 没有赋值的元素都为 0
var arrAge = [3]int{18, 20}
// 省略长度
var arrLazy = [...]int{5, 6, 7}
// k-v
var arryKeyValue = [5]string{3: "Chris", 4: "Ron"} // 只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串
```

### 7.1.3 多维数组

`[3][5]int`

### 7.1.4 将数组传递给函数

把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：

- 传递数组的指针
- 使用数组的切片

----

## 7.2 切片

- 切片是一个引用类型（像 Python 里的 list）
- `cap()` 测量切片最长能达到多少
- 因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用

```go
// 声明格式
var identifier []type

// 初始化格式 start ~ end-1
var identifier []type = arr1[start:end]

// 完整数组
var identifier []type = arr1[:]
var identifier = &arr1

// 类似数组的方式初始化（不指定长度）
var x = []int{2, 3, 4, 5}
```

!> 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针

### 7.2.2 将切片传递给函数

如果你有一个函数需要对数组做操作，你可能总是需要**把参数声明为切片**。当你调用该函数时，把数组分片，创建为一个 切片引用并传递给该函数。

### 7.2.3 用 make() 创建一个切片

当相关**数组还没有定义时**：

- 我们可以使用 `make()` 函数来创建一个切片
- 同时创建好相关数组：`var slice1 []type = make([]type, len)`

只占用 len 个数，而不占用整个数组：

```go
var slice2 []int = make([]int, len, cap)
```

所以下面两种方法可以生成相同的切片：

```go
make([]int, 50, 100)
new([100]int)[0:50]
```

### 7.2.4 new() 和 make() 的区别

- `new(T)`
  - 为每个新的类型 T 分配一片内存
  - 初始化为 0 并返回类型为 `*T` 的**内存地址**（指向类型为 T，值为 0 的地址的指针）
  - 适用于值类型（数组、结构体）
  - 相当于 `&T{}`
- `make(T)`
  - 返回一个类型为 T 的初始值
  - 适用于 3 种内建的引用类型：切片、map 和 channel

![](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/7.2_fig7.3.png?raw=true)

!> new 函数分配内存，make 函数初始化。

### 7.2.5 多维切片

和数组一样，切片通常也是一维的，但是也可以由一维组合成高维。

### 7.2.6 bytes 包

定义：

```go
var buffer bytes.Buffer
// 或使用 new 获得一个指针
var r *bytes.Buffer = new(bytes.Buffer)
// 或通过函数
func NewBuffer(buf []byte) *Buffer
```

通过 buffer 串联字符串：

> 这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候。

```go
buffer.WriteString(s) // 追加
buffer.String() // 转为 string
```

----

## 7.3 For-range 结构

```go
// value 只是 slice1 某个索引位置的值的一个拷贝
for ix, value := range slice {
    // do...
}
```

----

## 7.4 切片重组（reslice）

切片可以反复扩展直到占据整个相关数组。

```go
// 将切片扩展 1 位
sl = sl[0:len(sl)+1]
```

----

## 7.5 切片的复制与追加

1. 拷贝切片 `copy`
2. 向切片追加新元素 `append`

```go
sl_from := []int{1, 2, 3}
sl_to := make([]int, 10)
// copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst
n := copy(sl_to, sl_from) // 返回拷贝的元素个数 n = 3 

sl3 := []int{1, 2, 3}
// 将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片
// 如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储
sl3 = append(sl3, 4, 5, 6)
```

----

## 7.6 字符串、数组和切片的应用

### 7.6.1 从字符串生成字节切片

```go
// 获取一个字节的切片 c
c := []byte(s)
// 也可以用 copy
copy(dst []byte, src string)
```

将字符串追加到切片的末尾：

```go
var b []byte
var s string
b = append(b, s...)
```

### 7.6.2 获取字符串的某一部分

```go
substr := str[start:end] 
```

### 7.6.3 字符串和切片的内存结构

一个字符串实际上是一个双字结构：

1. 一个指向实际数据的指针
2. 记录字符串长度的整数

![](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/7.6_fig7.4.png)

### 7.6.4 修改字符串中的某个字符

!> Go 语言中的字符串是不可变的，也就是说 `str[index]` 这样的表达式是不可以被放在等号左侧的。

因此：

1. 先将字符串转换成字节数组
2. 通过修改数组中的元素值来达到修改字符串的目的
3. 将字节数组转换回字符串格式

```go
s := "hello"
c := []byte(s)
c[0] = '0'
s2 := string(c)
``` 

### 7.6.5 字节数组对比函数

```go
// 返回两个字节数组字典顺序的整数对比结果
func Compare(a, b[]byte) int {
    for i:=0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    // 数组的长度可能不同
    switch {
    case len(a) < len(b):
        return -1
    case len(a) > len(b):
        return 1
    }
    return 0 // 数组相等
}
```

### 7.6.6 搜索及排序切片和数组

`sort` 实现常见的搜索和排序操作。

- `func Ints(a []int)` 实现对 int 类型的切片排序 `sort.Ints(arri)`
- `IntsAreSorted(a []int) bool` 检查是否已经排过序
- `func Float64s(a []float64)` 排序 float64 元素
- `func Strings(a []string)` 排序字符串元素

### 7.6.7 append 函数常见操作

> 可以使用切片和 append 操作来表示任意可变长度的序列。

### 7.6.8 切片和垃圾回收

!> 只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。