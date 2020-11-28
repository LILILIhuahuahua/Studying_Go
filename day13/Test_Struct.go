package main

import "fmt"

func main(){
	p1:=person{"lixianhua"}
	fmt.Println(p1)

	changName(p1)

	fmt.Println(p1)

}

func changName(p person) {
	p.name = "黎先桦"
}
type person struct {
	name string
}
