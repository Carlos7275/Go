package main

import (
	"fmt"
	"packages/mylibrary"
)

func main() {
	var number float32
	var exp int

	fmt.Println("Turn a Number:")
	fmt.Scan(&number)
	fmt.Println("Turn a Exponential:")
	fmt.Scan(&exp)
	fmt.Println(mylibrary.Pow(number, exp))
}
