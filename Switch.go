package main

import (
	"fmt"
)

type data [2]int

func main() {
	switch x := 8; x {
	//没有匹配，就执行default
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough  //go的switch默认有break，fallthrough表强制继续执行
	case 6:
		x += 20
		fmt.Println(x)
	}

	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
}
