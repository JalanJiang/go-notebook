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