## 9.1 标准库概述

具有常用功能的内置包被称为标准库。

## 9.2 regexp 包

```go
searchIn := "Haaaaa"
pat := "[0-9]+.[0-9].+" // 正则

if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
    fmt.Println("Match Found")
}

// 更多方法中，必须先将正则模式通过 Compile 方法返回一个 Regexp 对象。然后我们将掌握一些匹配，查找，替换相关的功能
re, _ := regexp.Compile(pat)
str := re.ReplaceAllString(searchIn, "##.#") // 替换操作
```

## 9.3 锁和 sync 包

- map 类型是不存在锁的机制来实现这种效果(出于对性能的考虑)，所以 map 类型是非线程安全的。当并行访问一个共享的 map 类型的数据，map 数据将会出错。
- 锁的机制是通过 sync 包中 Mutex 来实现的。
- `sync.Mutex` 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
- 在 sync 包中还有一个 RWMutex 锁：他能通过 `RLock()` 来允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作。如果使用 `Lock()` 将和普通的 Mutex 作用相同。包中还有一个方便的 Once 类型变量的方法 `once.Do(call)`，这个方法确保被调用函数只能被调用一次。

```go
import "sync"

type Info struct {
    mu sync.Mutex
    // .. others
}

func Update(info *Info) {
    // 上锁
    info.mu.Lock()
    info.Str = "xxxx"
    // 释放锁
    info.mu.Unlock()
}
```

## 9.4 精密计算和 big 包

- 对于整数的高精度计算 Go 语言中提供了 big 包，被包含在 math 包下。

## 9.5 自定义包和可见性

- 当写自己包的时候，要使用短小的不含有 `_`（下划线）的小写单词来为文件命名。
- 主程序利用的包必须在主程序编写之前被编译。
- 按照惯例,子目录和包之间有着密切的联系：为了区分,不同包存放在不同的目录下，每个包(所有属于这个包中的 go 文件)都存放在和包名相同的子目录下。

### 导入包

import 的一般格式如下:

```
import "包的路径或 URL 地址" 
```

例如：

```go
import "github.com/org1/pack1”
```

Import with `.` :

```
import . "./pack1"
```

当使用.来做为包的别名时，你可以不通过包名来使用其中的项目。例如：`test := ReturnStr()`。在当前的命名空间导入 pack1 包，一般是为了具有更好的测试效果。

Import with `_` :

```go
import _ "./pack1/pack1"
```

pack1包只导入其副作用，也就是说，只执行它的 init 函数并初始化其中的全局变量。

#### 导入外部安装包

如果你要在你的应用中使用一个或多个外部包，首先你必须使用 `go install`（参见第 9.7 节）在你的本地机器上安装它们。

```
go install codesite.ext/author/goExample/goex
```

```go
import goex "codesite.ext/author/goExample/goex"
```

### 初始化

- 程序的执行开始于导入包，初始化 main 包然后调用 main 函数。
- 一个没有导入的包将通过分配初始值给所有的包级变量和调用源码中定义的包级 init 函数来初始化。

## 9.8 自定义包的目录结构、go install 和 go test

### 9.8.1 自定义包的目录结构

```
/home/user/goprograms
	ucmain.go	(uc包主程序)
	Makefile (ucmain的makefile)
	ucmain
	src/uc	 (包含uc包的go源码)
		uc.go
	 	uc_test.go
	 	Makefile (包的makefile)
	 	uc.a
	 	_obj
			uc.a
		_test
			uc.a
	bin		(包含最终的执行文件)
		ucmain
	pkg/linux_amd64
		uc.a	(包的目标文件)
```

### 9.8.2 本地安装包

```
go install /home/user/goprograms/src/uc # 编译安装uc
cd /home/user/goprograms/uc
go install ./uc 	# 编译安装uc（和之前的指令一样）
cd ..
go install .	# 编译安装ucmain
```
### 9.8.3 依赖系统的代码

???

## 9.9 通过 Git 打包和安装

### 9.9.1 安装到 GitHub

### 9.9.2 从 GitHub 安装

```
go get github.com/NNNN/uc
```

## 9.10 Go 的外部包和项目

[Go Walker](https://gowalker.org/) 支持根据包名在海量数据中查询。