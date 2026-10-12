package main

import (
	"fmt"
	"github.com/konsole-log/GO-STUFFS/structs_methods"
	"github.com/konsole-log/GO-STUFFS/integers"
)

func main() {
	sum := integers.Add(5,2)
	fmt.Println(sum)

	rectangle := structsmethods.Rectangle{10.0,20.0}
	answer:=rectangle.Area()
	fmt.Printf("%g",answer)
}
