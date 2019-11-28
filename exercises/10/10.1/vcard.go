/*
定义结构体 Address 和 VCard，后者包含一个人的名字、地址编号、出生日期和图像，试着选择正确的数据类型。构建一个自己的 vcard 并打印它的内容。

提示：
VCard 必须包含住址，它应该以值类型还是以指针类型放在 VCard 中呢？
第二种会好点，因为它占用内存少。包含一个名字和两个指向地址的指针的 Address 结构体可以使用 %v 打印：
{Kersschot 0x126d2b80 0x126d2be0}
*/
package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street      string
	HouseNumber uint32
	ZipCode     string
	City        string
	Country     string
}

type VCard struct {
	firstName string
	lastName  string
	BirthDate time.Time
	Addresses map[string]*Address
}

func main() {
	addr1 := &Address{"StreetName", 101, "", "Shenzhen", "China"}
	addr2 := &Address{"StreetName2", 201, "", "Fuzhou", "China"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birthdt := time.Date(1994, 11, 9, 0, 0, 0, 0, time.Local)
	vcard := &VCard{"Jalan", "Jiang", birthdt, addrs}
	fmt.Printf("VCard: %v\n", vcard)
	fmt.Printf("Address: %v %v", addr1, addr2)
}

/* Output:
VCard: &{Jalan Jiang 1994-11-09 00:00:00 +0800 CST map[now:0xc00009a050 youth:0xc00009a000]}
Address: &{StreetName 101  Shenzhen China} &{StreetName2 201  Fuzhou China}
*/
