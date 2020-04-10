package main

import (
	"fmt"
	"runtime"
)

var x int
var y string
var z bool

var m = 42
var n = "James Bond"
var o = true

type zeus int

var g zeus
var h int

func main() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Println("The value of x is:", a)
	fmt.Println("The value of y is:", b)
	fmt.Println("The value of z is:", c)

	fmt.Println("The value of x is:", x)
	fmt.Println("The value of y is:", y)
	fmt.Println("The value of z is:", z)

	s := fmt.Sprintf("%v\t%v\t%v\t", m, n, o)
	fmt.Println(s)

	fmt.Println("The value of g is:", g)
	fmt.Printf("%T\n", g)
	g = 42
	fmt.Println("The value of g is:", g)

	h = int(g)
	fmt.Println("The value of h is:", h)
	fmt.Printf("%T\n", h)

	// GOOS and GOARCH return the os and architechure spec for the pc
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
