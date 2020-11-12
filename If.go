package main

import (
	"fmt"
)

func main() {
	//if中同时初始化，num只能用于if中
	if num := 10; num % 2 == 0 { //checks if number is even
		fmt.Println(num,"is even")
	}  else {
		fmt.Println(num,"is odd")
	}
}