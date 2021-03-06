<!-- <p align="center"><img width="200px" src="https://blog.golang.org/lib/godoc/images/footer-gopher.jpg"></p> -->
<p align="center"><img width="200px" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTUIISmTbpsEq_63LdPlLk7WSgplC5Tq1NnzSPC_hcClWm0Uah7PQ
"></p>

# The Way To Go

- 记录 [《Go入门指南》](https://github.com/unknwon/the-way-to-go_ZH_CN) 学习笔记，[点此跳转目录](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)，[点此查看原仓库练习答案](https://github.com/unknwon/the-way-to-go_ZH_CN/tree/master/eBook/exercises)
- 使用 [docsify](https://docsify.js.org/#/) 构建，搭建教程见 [《docsify 入坑指南与我放弃 Gitbook 的那些理由》](http://jalan.space/2019/06/21/2019/begin-docsify/)

## 目录大纲

### 第一部分：学习 Go 语言

#### 第 1 章：Go 语言的起源，发展与普及

- [x] 1.1 起源与发展
- [ ] 1.2 语言的主要特性与发展的环境和影响因素

#### 第 2 章：安装与运行环境

- [ ] 2.1 平台与架构 
- [x] 2.2 Go 环境变量
- [ ] 2.3 在 Linux 上安装 Go
- [ ] 2.4 在 Mac OS X 上安装 Go
- [ ] 2.5 在 Windows 上安装 Go
- [ ] 2.6 安装目录清单
- [x] 2.7 Go 运行时（runtime）
- [x] 2.8 Go 解释器

#### 第 3 章：编辑器、集成开发环境与其它工具

- [ ] 3.1 Go 开发环境的基本要求
- [ ] 3.2 编辑器和集成开发环境
- [x] 3.3 调试器
- [x] 3.4 构建并运行 Go 程序
- [x] 3.5 格式化代码
- [x] 3.6 生成代码文档
- [ ] 3.7 其它工具
- [x] 3.8 Go 性能说明
- [ ] 3.9 与其它语言进行交互

### 第二部分：语言的核心结构与技术

#### 第 4 章：基本结构和基本数据类型

- [x] 4.1 文件名、关键字与标识符
- [x] 4.2 Go 程序的基本结构和要素
- [x] 4.3 常量
- [x] 4.4 变量
- [x] 4.5 基本类型和运算符
- [x] 4.6 字符串
- [x] 4.7 strings 和 strconv 包
- [x] 4.8 时间和日期
- [x] 4.9 指针

#### 第 5 章：控制结构

- [x] 5.1 if-else 结构
- [x] 5.2 测试多返回值函数的错误
- [x] 5.3 switch 结构
- [x] 5.4 for 结构
- [x] 5.5 Break 与 continue
- [x] 5.6 标签与 goto

#### 第 6 章：函数（function）

- [x] 6.1 介绍
- [x] 6.2 函数参数与返回值
- [x] 6.3 传递变长参数
- [x] 6.4 defer 和追踪
- [x] 6.5 内置函数
- [x] 6.6 递归函数
- [x] 6.7 将函数作为参数
- [x] 6.8 闭包
- [x] 6.9 应用闭包：将函数作为返回值
- [x] 6.10 使用闭包调试
- [x] 6.11 计算函数执行时间
- [x] 6.12 通过内存缓存来提升性能

#### 第 7 章：数组与切片

- [x] 7.1 声明和初始化
- [x] 7.2 切片
- [x] 7.3 For-range 结构
- [x] 7.4 切片重组（reslice）
- [x] 7.5 切片的复制与追加
- [x] 7.6 字符串、数组和切片的应用

#### 第 8 章：Map

- [x] 8.1 声明、初始化和 make
- [x] 8.2 测试键值对是否存在及删除元素
- [x] 8.3 for-range 的配套用法
- [x] 8.4 map 类型的切片
- [x] 8.5 map 的排序
- [x] 8.6 将 map 的键值对调

#### 第9章：包（package）

- [x] 9.1 标准库概述
- [x] 9.2 regexp 包
- [x] 9.3 锁和 sync 包
- [ ] 9.4 精密计算和 big 包
- [x] 9.5 自定义包和可见性
- [ ] 9.6 为自定义包使用 godoc
- [x] 9.7 使用 go install 安装自定义包
- [x] 9.8 自定义包的目录结构、go install 和 go test
- [x] 9.9 通过 Git 打包和安装
- [x] 9.10 Go 的外部包和项目
- [ ] 9.11 在 Go 程序中使用外部库

#### 第10章：结构（struct）与方法（method）

- [x] 10.1 结构体定义
- [x] 10.2 使用工厂方法创建结构体实例
- [x] 10.3 使用自定义包中的结构体
- [x] 10.4 带标签的结构体
- [x] 10.5 匿名字段和内嵌结构体
- [x] 10.6 方法
- [x] 10.7 类型的 String() 方法和格式化描述符
- [x] 10.8 垃圾回收和 SetFinalizer

#### 第11章：接口（interface）与反射（reflection）

- [x] 11.1 接口是什么
- [x] 11.2 接口嵌套接口
- [x] 11.3 类型断言：如何检测和转换接口变量的类型
- [x] 11.4 类型判断：type-switch
- [x] 11.5 测试一个值是否实现了某个接口
- [x] 11.6 使用方法集与接口
- [x] 11.7 第一个例子：使用 Sorter 接口排序
- [x] 11.8 第二个例子：读和写
- [x] 11.9 空接口
- [x] 11.10 反射包
- [x] 11.11 Printf 和反射
- [x] 11.12 接口与动态类型
- [x] 11.13 总结：Go 中的面向对象
- [x] 11.14 结构体、集合和高阶函数

### 第三部分：Go 高级编程

#### 第12章：读写数据

- [ ] 12.1 读取用户的输入
- [ ] 12.2 文件读写
- [ ] 12.3 文件拷贝
- [ ] 12.4 从命令行读取参数
- [ ] 12.5 用 buffer 读取文件
- [ ] 12.6 用切片读写文件
- [ ] 12.7 用 defer 关闭文件
- [ ] 12.8 使用接口的实际例子：fmt.Fprintf
- [ ] 12.9 格式化 JSON 数据
- [ ] 12.10 XML 数据格式
- [ ] 12.11 用 Gob 传输数据
- [ ] 12.12 Go 中的密码学

#### 第13章：错误处理与测试

- [x] 13.1 错误处理
- [x] 13.2 运行时异常和 panic
- [x] 13.3 从 panic 中恢复（Recover）
- [x] 13.4 自定义包中的错误处理和 panicking
- [x] 13.5 一种用闭包处理错误的模式
- [x] 13.6 启动外部命令和程序
- [ ] 13.7 Go 中的单元测试和基准测试
- [ ] 13.8 测试的具体例子
- [ ] 13.9 用（测试数据）表驱动测试
- [ ] 13.10 性能调试：分析并优化 Go 程序