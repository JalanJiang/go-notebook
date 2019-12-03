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
// 返回指针
t = new(T)
```

### 初始化

```go
// 返回指针
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

## 10.2 使用工厂方法创建结构体实例

### 10.2.1 结构体工厂

- 工厂名字以 `new` 或 `New` 开头
- 返回一个指向结构体实例的指针

```go
func newFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }

    return &File{fd, name}
}
```

这样调用：

```go
f := NewFile(10, "./test.txt")
```

与其他语言中 `File f = new File(...)` 等价：

1. `new(File)`
2. `&File{}`

### 10.2.2 map 和 struct vs new() 和 make()

可以使用 `make` 的内置类型：

- slices
- maps
- channels

## 10.4 带标签的结构体

结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）。

- 附属于字段的字符串
- **只有包 `reflect` 能获取它**

```go
package main

import (
    "fmt"
    "reflect"
)

type TagType struct { // tags
    field1 bool "An important answer"
    field2 string "The name of the thing"
    field3 int "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix) // 取出第 ix 个字段
	fmt.Printf("%v\n", ixField.Tag)
}

/* Output:
An important answer
The name of the thing
How much there are
*/
```

## 10.5 匿名字段和内嵌结构体

### 10.5.1 定义

结构体可以包含一个或多个 匿名（或内嵌）字段

- 这些字段没有显式的名字
- 只有字段的类型是必须的
- 此时类型就是字段的名字

**可以粗略地将这个和面向对象语言中的继承概念相比较**。Go 语言中的继承是通过**内嵌**或**组合**来实现的，所以可以说，在 Go 语言中，**相比较于继承，组合更受青睐**。

```go
package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // 匿名字段
	innerS // 匿名结构体
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

/*Output:
outer.b is: 6
outer.c is: 7.500000
outer.int is: 60
outer.in1 is: 5
outer.in2 is: 10
outer2 is:{6 7.5 60 {5 10}}
*/
```

### 10.5.2 内嵌结构体

这个简单的“继承”机制提供了一种方式，使得可以从另外一个或一些类型继承部分或全部实现。

```go
package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func main() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}

/*Output:
1 2 3 4
{1 2}
*/
```

### 10.5.3 命名冲突

1. **外层名字会覆盖内层名字**（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。

## 10.6 方法

### 10.6.1 方法是什么

**在 Go 语言中，结构体就像是类的一种简化形式。**

- Go 方法是作用在接收者（receiver）上的一个函数，**接收者是某种类型的变量**（很像对象）
- 方法是一种特殊类型的函数
- 类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集（method set）

!> 任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。但是接收者不能是一个接口类型（参考 第 11 章），因为接口是一个抽象定义，但是方法却是具体实现；如果这样做会引发一个编译错误：invalid receiver type…

- 一个类型加上它的方法等价于面向对象中的一个类
- 类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件
- 唯一的要求是：它们必须是同一个包的

#### 定义方法的格式：

```go
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
```

- 如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 `object.name` 选择器符号：`recv.Method1()`
- 如果方法不需要使用 `recv` 的值，可以用 `_` 替换它

!> recv 就像是面向对象语言中的 `this` 或 `self`，但是 Go 中并没有这两个关键字。随个人喜好，你可以使用 `this` 或 `self` 作为 receiver 的名字。

```go
package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

/*Output:
The sum is: 22
Add them to the param: 42
The sum is: 7
*/
```

### 10.6.2 函数和方法的区别

- 函数将变量作为参数：Function1(recv)
- 方法在变量上被调用：recv.Method1()

!> 方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

### 10.6.3 指针或值作为接收者

如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。否则，就在普通的值类型上定义方法。

!> 指针方法和值方法都可以在指针或非指针上被调用。

### 10.6.4 方法和未导出字段

获取未导出字段：使用 getter 和 setter 方法。

- 对于 setter 方法使用 Set 前缀
- 对于 getter 方法只使用成员名

```go
package person

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
```

使用：

```go
package main

import (
	"./person"
	"fmt"
)

func main() {
	p := new(person.Person)
	// p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	// p.firstName = "Eric"
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName()) // Output: Eric
}
```

!> 对象的字段（属性）不应该由 2 个或 2 个以上的不同线程在同一时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用包 sync（参考第 9.3 节）中的方法。在第 14.17 节中我们会通过 goroutines 和 channels 探索另一种方式。

### 10.6.5 内嵌类型的方法和继承

当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型**继承**了这些方法：**将父类型放在子类型中来实现亚型**。

### 10.6.6 如何在类型中嵌入功能

### 10.6.7 多重继承

### 10.6.8 通用方法和方法命名

### 10.6.9 和其他面向对象语言比较 Go 的类型和方法


