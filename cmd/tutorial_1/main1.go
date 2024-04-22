package main

import "fmt"

func main() {
	var intNum uint16 = 32767
	fmt.Println(intNum)

	var floatNum float64
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	//same datatype can only perform arithmetic operation
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
}
