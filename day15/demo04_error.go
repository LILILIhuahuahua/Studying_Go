package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	files,err:=filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern{
		fmt.Println(err) //syntax error in pattern
		return
	}
	fmt.Println("files:",files)
}
