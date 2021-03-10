package main

import(
	"fmt"
	lib "github.com/nagoyan55/go-firstlib/lib"
)

func main(){
	fmt.Println("Hello, World")
	x := lib.Add(3, 4)
	fmt.Println(x)
}
