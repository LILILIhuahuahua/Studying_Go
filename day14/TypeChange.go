package main

import (
	"fmt"
)

func main() {

	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1)
	//i1 = i2 //cannot use i2 (type int) as type myint in assignment
	fmt.Println(i1,i2)
	var i3 myint2
	i3 = i2
	fmt.Println(i1,i2,i3)

}

type myint int
type myint2 = int //不是重新定义类型，只是给int起别名
