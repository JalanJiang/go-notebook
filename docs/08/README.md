## 8.1 声明、初始化和 make

### 8.1.1 概念

- map 是引用类型
- 声明时不需要知道 map 长度，map 动态增长
- 通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比**从数组和切片的索引中直接读取**要慢 100 倍

**声明：**

```go
// var map1 map[keytype]valuetype
var map1 map[string]int
```

**初始化：**

```go
mapList = map[string]int{"one": 1, "two": 2}
```

**分配内存：**

- map 是 引用类型 的： 内存用 `make` 方法来分配
- 不要使用 `new`，永远用 `make` 来构造 map

```go
var map1 map[string]int // 声明
var map1 = make(map[keytype]valuetype) // 分配内存
// 或者额可以直接
map2 := make(map[keytype]valuetype)
```

!> 如果你错误的使用 `new()` 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址。

**赋值：**

```go
map1[key1] = valu1
```

### 8.1.2 map 容量

- 和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制
- 但也可以选择标明 map 的初始容量 `capacity`
  - 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1
  - **对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明**

```go
map3 := make(map[int]string, 100)
```

### 8.1.3 用切片作为 map 的值

一个 key 对应多个 value 的解决方法：

```go
// 将 value 定义为 []int 型
map1 := make(map[int][]int)
map2 := make(map[int]*[]int)
```

## 8.2 测试键值对是否存在及删除元素

```go
val1, isPresent = map1[key1]
```

和 if 混用：

```go
if _, ok := map1[key1]; ok {
    // ...
}
```

从 map1 删除 key1:

```go
// 如果 key1 不存在，该操作不会产生错误
delete(map1, key1)
```

---- 

## 8.3 for-range 的配套用法

```go
for key, value := range map1 {
    // ...
}
```

----

## 8.4 map 类型的切片

假设我们想获取一个 map 类型的切片，我们必须使用两次 `make()` 函数：

1. 分配切片
2. 分配切片中每个 map 元素

```go
items := make([]map[int]int, 5)
for i := range items {
    items[i] = make(map[int]int, 1)
    items[i][1] = 2
}
```

----

## 8.5 map 的排序

!> map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序。

如果你想为 map 排序：

1. 将 key（或者 value）拷贝到一个切片
2. 对切片排序
3. 使用切片的 for-range 方法打印出所有的 key 和 value

如果你想要一个排序的列表你最好使用结构体切片，这样会更有效：

```go
type name struct {
    key string
    value int
}
```

----

## 8.6 将 map 的键值对调

1. 创建一个新的 map
2. 循环，对新的 map 赋值