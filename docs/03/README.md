[编辑器和 IDE 扩展](http://go-lang.cat-v.org/text-editors/)

## 3.4 构建并运行 Go 程序

- `go build` 编译自身包和依赖包
- `go install` 编译并安装自身包和依赖包

## 3.5 格式化代码

必须在编译或提交版本管理系统之前使用 `gofmt` 来格式化你的代码。

## 3.6 生成代码文档

`go doc` 工具会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。

用法：

- `go doc package` 获取包的文档注释，例如：`go doc fmt` 会显示使用 `godoc` 生成的 `fmt` 包的文档注释。
- `go doc package/subpackage` 获取子包的文档注释，例如：`go doc container/list`
- `go doc package function` 获取某个函数在某个包中的文档注释，例如：`go doc fmt Printf` 会显示有关 `fmt.Printf()` 的使用说明。

启动本地服务器：`godoc -http=:6060`