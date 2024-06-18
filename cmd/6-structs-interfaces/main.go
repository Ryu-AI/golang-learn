package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner // can also just be set to owner type instead of ownerInfo
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// type owner struct {
// 	name string
// }

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	// var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name) //this would become myEngine.name

	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)

	//anonymous struct -> not reusable
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{35, 25}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	canMakeIt(myEngine, 50)
}
