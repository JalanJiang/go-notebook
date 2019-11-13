package main

var a string

func main() {
	a = "G" // 在函数体外申明：全局变量
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}

// 输出：GOG