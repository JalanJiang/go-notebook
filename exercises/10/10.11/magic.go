package main

import (
	"fmt"
)

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	// 调用的是 Voodoo 的 Magic
	v.Magic()
	// 调用的是 Base 的 MoreMagic 和 Magic
	v.MoreMagic()
}

/*Output:
voodoo magic
base magic
base magic
*/
