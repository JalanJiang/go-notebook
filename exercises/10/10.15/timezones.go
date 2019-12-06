package main

import "fmt"

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

// 定义 map
var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time"}

func (tz TZ) String() string {
	if zone, ok := timeZones[tz]; ok {
		return zone
	}
	return ""
}

func main() {
	fmt.Println(EST) // Print* knows about method String() of type TZ
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}

/*
Eastern Standard time
Universal Greenwich time
Central Standard time
*/
