package main

import "fmt"

func main(){
	//创建数组，值传递
	var array = [...]int{1,2,3,4}
	var array2 = array
	array2[0] = 0
	//array没改，就是值传递
	fmt.Print(array,array2)
	fmt.Print(len(array),cap(array))

	fmt.Println()
	//创建切片，地址传递
	var slice = []int{1,2,3,4}
	var slice2 = slice
	slice2[0] = 0
	//slice改，就是地址传递
	fmt.Print(slice,slice2)
	fmt.Print(len(array),cap(array2))
}
