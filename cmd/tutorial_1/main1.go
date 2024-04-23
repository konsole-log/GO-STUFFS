package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// method of gasEngine
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// method of electricEngine
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// to show use of interface
type engine interface {
	milesLeft() uint8
}

// now this function can take both engine types
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank:%v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 10) //due to interface we can use different engine type

	var eEngine electricEngine = electricEngine{30, 30}
	fmt.Println(eEngine.kwh, eEngine.mpkwh)
	fmt.Printf("Total miles left in battery:%v\n", eEngine.milesLeft())
	canMakeIt(eEngine, 20) //due to interface we can use different engine type
}
