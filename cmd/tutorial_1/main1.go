package main

import (
	"errors"
	"fmt"
)

func main() {
	// var intNum uint16 = 32767
	// fmt.Println(intNum)

	// var floatNum float64
	// fmt.Println(floatNum)

	// var floatNum32 float32 = 10.1
	// var intNum32 int32 = 2
	// //same datatype can only perform arithmetic operation
	// var result float32 = floatNum32 + float32(intNum32)
	// fmt.Println(result)

	// var myString string = "hello world"
	// fmt.Println(myString)

	// var name string = `this
	// is formatted right away`
	// fmt.Println(name)

	// var name1 string = "this " + " is concatenated"
	// fmt.Println(name1)
	// ///gives bytes of the word test go uses utf8
	// fmt.Println(len("test"))

	// fmt.Println(utf8.RuneCountInString("test"))

	// var myRune rune = 'a' //like character?
	// fmt.Println(myRune)

	// var myBoolean bool = false
	// fmt.Println(myBoolean)

	// var mystr string
	// fmt.Println(mystr) //'' defautl value for string

	// var mynum int32
	// fmt.Println(mynum) //0 for float,int,rune
	// //we can declare variable following ways:
	// var fid = "hello"
	// var myNum = 1
	// greetings := "yoo"
	// a := 2
	// fmt.Println(myNum)
	// fmt.Println(fid)
	// fmt.Println(greetings)
	// fmt.Println(a)
	// //applies same fule as variables
	// const myConst string = "constvalue"
	// fmt.Println(myConst)
	// const pi float32 = 3.1415
	// fmt.Println(pi)

	printMe("hello")

	var numerator int = 11
	var denominator int = 2
	var result, rem, err = intDivision(numerator, denominator)
	//if else statement
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if rem == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("quotioent:%v and remainder:%v", result, rem)

	// }
	//normal switch
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("quotioent:%v and remainder:%v", result, rem)
	}

	//conditional switch statement
	switch rem {
	case 0:
		fmt.Printf("THe division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

// func functionName(parameters)(returning dattypes)
func intDivision(numerator int, dennominator int) (int, int, error) {
	var err error
	if dennominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / dennominator
	var rem int = numerator % dennominator
	return result, rem, err
}
