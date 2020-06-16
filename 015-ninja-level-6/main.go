package main

import (
	"fmt"
	"math"
)

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "is the answer to everything"
}

// function using variadic parameters
func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

// function using a slice
func xsum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

// attaching a struct to a method
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}

// using interfaces
type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

// returning a function
func ret() func() int {
	return func() int {
		return 42
	}
}

// passing a function into a fucntion as a parameter (CALLBACK)
func funcParm(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

// an example of closure
func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {

	defer func() {
		fmt.Println("this is printed last!")
	}()

	fmt.Println(foo())
	fmt.Println(bar())
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(xi...))
	fmt.Println(xsum(xi))

	// attaching a method to a struct
	p1 := person{
		first: "John",
		last:  "Wick",
		age:   32,
	}
	// call ing the method on the struct
	p1.speak()

	// using interfaces
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)

	// annoymus function
	func() {
		fmt.Println("This is an annoymus function")
	}()

	// creating a function and assigning it to a variable
	// then calling that function
	x := func(j int) int {
		sum := 0
		for i := 0; i < j; i++ {
			sum += i
		}
		return sum
	}

	fmt.Printf("type for x: %T\n", x)
	fmt.Println("value for x(5): ", x(5))

	// returning a function
	f := ret()
	fmt.Println(f())

	// passing a function into a fucntion as a parameter (CALLBACK)
	gi := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	fi := funcParm(gi, []int{1, 2, 3, 4, 5, 10})
	fmt.Println(fi)

	// an example of closure
	cf1 := closure()
	fmt.Println(cf1())
	fmt.Println(cf1())
	cf2 := closure()
	fmt.Println(cf2())
	fmt.Println(cf2())
	fmt.Println(cf2())
	fmt.Println(cf2())
	fmt.Println(cf2())
	fmt.Println(cf2())
}
