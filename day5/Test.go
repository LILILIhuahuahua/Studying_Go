package main

import "fmt"

func main(){
	var a,b,c int
	for i:=100;i<1000;i++{
		temp:=i
		a = temp/100
		temp%=100
		b = temp/10
		temp%=10
		c=temp
		if a*a*a* +b*b*b+c*c*c == i{
			fmt.Printf("%v是水仙花数",i)
		}
	}
}
