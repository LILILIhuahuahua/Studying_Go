package main

import "fmt"
func main(){
	var array =[]int{0,1,2,3,4,5,6,7,8,9}

	a1:=array[2:4]
	fmt.Print("a1的cap为",cap(a1))
}
