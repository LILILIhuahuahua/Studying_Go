package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入您的成绩：")
	fmt.Scanf("%d",&score)

	if(score>60){
		fmt.Printf("你的成绩及格")
	}else{
		fmt.Printf("不及格")
	}
}