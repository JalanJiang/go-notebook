## 4.1 文件名、关键字与标识符

- `_`  被称为空白标识符
- 每个语句不需要像 C 家族中的其它语言一样以分号 `;` 结尾

----

## 4.2 Go 程序的基本结构和要素

```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world")
}
```

### 4.2.1 包的概念、导入与可见性

**包的概念：**

- 每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容
- 每个 Go 文件都属于且仅属于一个包
- 一个包可以由许多以 `.go` 为扩展名的源文件组成
- 必须在源文件中非注释的第一行指明这个文件属于哪个包
- `package main` 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 `main` 的包
- 所有的包名都应该使用小写字母

**编译顺序：**

- 如果想要构建一个程序，则包和包内的文件都必须以正确的顺序进行编译。包的依赖关系决定了其构建顺序
- 属于同一个包的源文件必须全部被一起编译，一个包即是编译时的一个单元，因此根据惯例，每个目录都只包含一个包
- 每一段代码只会被编译一次
- Go 中的包模型采用了显式依赖关系的机制来达到快速编译的目的

```
如果 A.go 依赖 B.go，而 B.go 又依赖 C.go：

编译 C.go, B.go, 然后是 A.go.
为了编译 A.go, 编译器读取的是 B.o 而不是 C.o.
```

**导入包：**

- 通过 `import` 关键字将一组包链接在一起

```go
// 因式分解关键字
import (
   "fmt"
   "os"
)
```

对包名重新设置：

```go
package main

import fm "fmt"

func main() {
  fm.Println("hello world")
}
```

**可见性规则：**

- 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个**大写字母**开头，如：Group1，那么使用这种形式的标识符的对象就**可以被外部包的代码所使用** - 类似 public
- 标识符如果以小写字母开头，则对包外是不可见的 - 类似 private

### 4.2.2 函数

```go
func functionName(parameter_list) (return_value_list) {
}
```

- 每一个可执行程序必须包含 `main` 函数：**既没有参数，也没有返回类型**
- 程序正常退出的代码为 0 即 `Program exited with code 0`

#### 4.2.3 注释

- 在多段注释之间应以空行分隔加以区分

**包注释：**

```go
// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman
```

**函数注释：**

- 出现在函数之前则要以函数名开头写注释

```go
// enterOrbit causes Superman to fly into low Earth orbit, a position
// that presents several possibilities for planet salvation.
func enterOrbit() error {
   ...
}
```

### 4.2.4 类型

- 使用 var 声明的变量的值会自动初始化为该类型的零值

**函数作为返回值：**

```go
func FunctionName (a typea, b typeb) typeFunc
```

**函数多返回值：**

```go
func FunctionName (a typea, b typeb) (t1 type1, t2 type2)

// 返回形式
return var1, var2
```

**使用 `type` 自定义类型：**

```go
type IZ int
var a IZ = 5
```

因式分解关键字的方法：

```go
type (
   IZ int
   FZ float64
   STR string
)
```

### 4.2.5 Go 程序的一般结构

- 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
- 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
- 如果当前包是 main 包，则定义 main 函数。
- 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。

```go
package main

import (
   "fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
   var a int
   Func1()
   // ...
   fmt.Println(a)
}

func (t T) Method1() {
   //...
}

func Func1() { // exported function Func1
   //...
}
```

**执行顺序：**

按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：

1. 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
2. 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
3. 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

#### 4.2.6 类型转换

- 所有的转换都必须显式说明

```go
valueOfTypeB = typeB(valueOfTypeA)
```

例如：

```go
a := 5.0
b := int(a)
```

**具有相同底层类型的变量之间可以相互转换：**

```go
var a IZ = 5
c := int(a)
d := IZ(c)
```

### 4.2.7 Go 命名规范

- 名称不需要指出自己所属的包，因为在调用的时候会使用包名作为限定符
- 返回某个对象的函数或方法的名称一般都是使用**名词**，没有 `Get...` 之类的字符，如果是用于修改某个对象，则使用 `SetName`
- 可以使用大小写混合的方式，如 `MixedCaps` 或 `mixedCaps`，而不是使用下划线来分割多个名称。

----

## 4.3 常量

- 使用关键字 `const` 定义，用于存储不会改变的数据
- 存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
- `const identifier [type] = value`
- 未定义类型的常量会在必要时刻根据上下文来获得相关类型
- 常量的值必须是能够在编译时就能够确定的
- 数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出
- 反斜杠 `\` 可以在常量表达式中作为多行的连接符使用

**用作枚举：**

```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

**`iota` 枚举值：**

- 每当 `iota` 在新的一行被使用时，它的值都会自动加 1
- 每遇到一次 const 关键字，`iota` 就重置为 0 

```go
const (
    a = iota
    b = iota
    c = iota
)

// 也可以这样简写
const (
    a = iota
    b
    c
)
```

## 4.4 变量

### 4.4.1 简介

- 声明变量的一般形式是使用 `var` 关键字：`var identifier type`
- 一个变量被声明之后，系统自动赋予它该类型的零值。**所有的内存在 Go 中都是经过初始化的**
- 变量的命名规则遵循**骆驼命名法**


```go
var a int
var b bool
var str string
```

**因式分解关键字写法：**

- 一般用于声明全局变量

```go
var (
    a int
    b bool
    str string
)
```

**变量声明+赋值：**

```go
var a int = 15
// Go 可以在编译时完成类型推断
var a = 15
```

**简短声明语法：**

- 当你在函数体内声明局部变量时，应使用简短声明语法 `:=`

```go
a := 1
```

### 4.4.2 值类型和引用类型

- 内存：具有 32 位（4 字节）或 64 位（8 字节）的长度
- 值类型：内存存值
- 引用类型：内存存指向值的地址

Go 中的引用类型：指针、slices、maps、channel

### 4.4.3 打印

fmt 包中的 `Printf`

### 4.4.4 简短形式，使用 := 赋值操作符

- **只能被用在函数体内，而不可以用于全局变量的声明与赋值**
- 在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明
- 全局变量是允许声明但不使用

### 4.4.5 init 函数

- 不能够被人为调用
- 在每个包完成初始化后自动执行
- 执行优先级比 main 函数高
- 用途：在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性

----

## 4.5 基本类型和运算符

- 只有两个类型相同的值才可以进行比较

### 4.5.1 布尔类型 bool

- 在格式化输出时，你可以使用 `%t` 来表示你要输出的值为布尔型

**命名**：函数以 `is` 或 `Is` 开头，例如 `isSorted`

### 4.5.2 数字类型

#### 4.5.2.1 整型 int 和浮点型 float

- Go 语言中没有 float 类型。（Go语言中只有 float32 和 float64）没有double类型
- 应该尽可能地使用 float64，因为 `math` 包中所有有关数学运算的函数都会要求接收这个类型
- Go 中不允许不同类型之间的混合使用，但是对于常量的类型限制非常少，因此允许常量之间的混合使用

```go
package main

func main() {
        var a int
        var b int32
        b = a + a // 编译错误
        b = b + 5 // 因为 5 是常量，所以可以通过编译
}
```

整数：

- int8（-128 -> 127）
- int16（-32768 -> 32767）
- int32（-2,147,483,648 -> 2,147,483,647）
- int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）

无符号整数：

- uint8（0 -> 255）
- uint16（0 -> 65,535）
- uint32（0 -> 4,294,967,295）
- uint64（0 -> 18,446,744,073,709,551,615）

浮点型（IEEE-754 标准）：

- float32（+- 1e-45 -> +- 3.4 * 1e38）
- float64（+- 5 * 1e-324 -> 107 * 1e308）

#### 4.5.2.2 复数

> todo

#### 4.5.2.3 位运算

- 与 &
- 或 |
- 异或 ^
- 位清除 &^ -> 0
- 按位补足 `^` ???
- 左移 `<<`
- 右移 `>>`

#### 4.5.2.4 逻辑运算符

```
==
!=
<
>
<=
>=
```

#### 4.5.2.5 算术运算符

```
+
-
*
/
%
```

#### 4.5.2.6 随机数

`rand` 包实现随机数。

random.go：

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8) // 返回介于 [0, n) 之间的伪随机数
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	times := int64(time.Now().Nanosecond())
	rand.Seed(times) // 使用 Seed(value) 函数来提供伪随机数的生成种子
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32()) // 返回介于 [0.0, 1.0) 之间的伪随机数
	}
}
```

### 4.5.3 运算符与优先级

```
优先级 	运算符
 7 		^ !
 6 		* / % << >> & &^
 5 		+ - | ^
 4 		== != < <= >= >
 3 		<-
 2 		&&
 1 		||
```

### 4.5.4 类型别名

TZ 就是 int 类型的新名称（用于表示程序中的时区）。

`type TZ int`

### 4.5.5 字符类型

- 字符不是 Go 语言的一个类型，**字符只是整数的特殊用例**
  
`byte` 类型是 `uint8` 的别名。

```
var ch byte = 65
var ch byte = '\x41' // \x 总是紧跟着长度为 2 的 16 进制数
var ch byte = '\377' // \ 后面紧跟着长度为 3 的八进制数
```

- Unicode 字符在内存中**使用 int 来表示**
- 在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 `\u` 或者 `\U`
  - `\u` 则总是紧跟着长度为 4 的 16 进制数
  - `\U` 紧跟着长度为 8 的 16 进制数
- 因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示
- 格式化说明符 `%c` 用于表示字符

包 unicode 包含的函数（其中 ch 代表字符）：

- 判断是否为字母：`unicode.IsLetter(ch)`
- 判断是否为数字：`unicode.IsDigit(ch)`
- 判断是否为空白符号：`unicode.IsSpace(ch)`

## 4.6 字符串

- Go 中的字符串里面的字符也可能根据需要占用 1 至 4 个字节
  - 这与其它语言如 C++、Java 或者 Python 不同（Java 始终使用 2 个字节）
  - 好处是不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码
- 可以通过索引获取字符
- 获取字符串中某个字节的地址的行为是非法的

**2 种形式的字面值：**

- 解释字符串：双引号
- 非解释字符串：反引号

**字符串拼接方法：**

1. `+`
2. `strings.Join()` 
3. 字节缓冲（`bytes.Buffer`）拼接

### 练习

> 创建一个用于统计字节和字符（rune）的程序，并对字符串 `asSASA ddd dsjkdsjs dk` 进行分析，然后再分析 `asSASA ddd dsjkdsjsこん dk`，最后解释两者不同的原因（提示：使用 unicode/utf8 包）。

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str2))
}
```

输出：

```
API server listening at: 127.0.0.1:30602
The number of bytes in string str1 is 22
The number of characters in string str1 is 22
The number of bytes in string str2 is 28
The number of characters in string str1 is 24
```

----

## 4.7 strings 和 strconv 包

Go 中使用 `strings` 包来完成对字符串的主要操作。

### 4.7.1 前缀和后缀

```go
// HasPrefix 判断字符串 s 是否以 prefix 开头
strings.HasPrefix(s, prefix string) bool
// HasSuffix 判断字符串 s 是否以 suffix 结尾
strings.HasSuffix(s, suffix string) bool
```

### 4.7.2 字符串包含关系

```go
// Contains 判断字符串 s 是否包含 substr
strings.Constains(s, substr string) bool
```

### 4.7.3 判断子字符串或字符在父字符串中出现的位置（索引）

```go
// Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引）
// -1 表示字符串 s 不包含字符串 str
strings.Index(s, str string) int
// LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引）
// 1 表示字符串 s 不包含字符串 str
strings.LastIndex(s, str string) int
// 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位
strings.IndexRune(s string, r rune) int
```

### 4.7.4 字符串替换

```go
// Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串
// n = -1 则替换所有字符串 old 为字符串 new
strings.Replace(str, old, new, n) string
```

### 4.7.5 统计字符串出现次数

```go
// Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数
strings.Count(s, str string) int
```

### 4.7.6 重复字符串

```go
// Repeat 用于重复 count 次字符串 s 并返回一个新的字符串
strings.Repeat(s, count int) string
```

### 4.7.7 修改字符串大小写

```go
strings.ToLower(s) string
strings.ToUpper(s) string
```

### 4.7.8 修剪字符串

```go
// 剔除字符串开头和结尾的空白符号
strings.TrimSpace(s)
// 剔除指定字符
strings.Trim(s, "cut") // 剔除头尾的 cut

strings.TrimLeft(s)
strings.TrimRight(s)
```

### 4.7.9 分割字符串

```go
// 会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice
strings.Fields(s)
// 用于自定义分割符号来对指定字符串进行分割，同样返回 slice
strings.Split(s, sep)
```

### 4.7.10 拼接 slice 到字符串

```go
// Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
strings.Join(sl, []string, sep string) string
```

- [官方文档](https://golang.org/pkg/strings/)
- [国内用户可访问](http://docs.studygolang.com/pkg/strings/)

### 4.7.11 从字符串中读取内容

```go
// 用于生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
strings.NewReader(str)
```

### 4.7.12 字符串与其它类型的转换

通过 `strconv` 包实现。

**数字 -> 字符串：**

```go
// 返回数字 i 所表示的字符串类型的十进制数
strconv.Itoa(i int) string
// 将 64 位浮点型的数字转换为字符串
strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
```

- 其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'）
- prec 表示精度
- bitSize 则使用 32 表示 float32，用 64 表示 float64

**字符串 -> 数字：**

```go
// 将字符串转换为 int 型
strconv.Atoi(s string) (i int, err error)
// 将字符串转换为 float64 型
strconv.ParseFloat(s string, bitSize int) (f float64, err error)
```

----

## 4.8 时间和日期

- `func (t Time) Format(layout string) string`，可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串

```go
// 当前时间
t := time.Now()
fmt.Printf("%02d.%02d.%4d", t.Day(), t.Month(), t.Year())
// UTC
t  = time.Now().UTC()
// 计算
week = 60 * 60 * 24 * 7 * 1e9
week_from_now := t.Add(time.Duration(week))
// formatting
fmt.Println(t.Format(time.RFC822))
```

----

## 4.9 指针

- 程序在内存中存储它的值，每个内存块（或字）有一个地址，通常用十六进制数表示
- 地址可以存储在一个叫做指针的特殊数据类型中
- 一个指针变量可以指向任何一个值的内存地址
- 可以在指针类型前面加上 `*` 号（前缀）来获取指针所指向的内容，这里的 `*` 号是一个类型更改器。使用一个指针引用一个值被称为间接引用（反引用）
- 对一个空指针的反向引用是不合法的，并且会使程序崩溃

```go
s := "hi"
// 声明
var p *string = &s
*p = "hello"
```

> 指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。指针传递是很廉价的，只占用 4 个或 8 个字节。当程序在工作中需要占用大量的内存，或很多变量，或者两者都有，使用指针会减少内存占用和提高效率。被指向的变量也保存在内存中，直到没有任何指针指向它们，所以从它们被创建开始就具有相互独立的生命周期。