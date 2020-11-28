package main

import (
	"fmt"
	"strconv"
)

func main() {

	res1 := fun1()
	fmt.Println(res1(10,20))
}

//函数类型
type my_fun  func (int,int)(string)

//fun1()函数，无参数，返回值是my_func类型
func fun1 () my_fun{
	fun := func(a,b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}