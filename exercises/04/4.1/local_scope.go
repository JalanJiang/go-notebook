package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	a := "O" // 修改的是局部变量
	print(a)
}

// 输出 GOG
