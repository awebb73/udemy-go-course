package main

import "fmt"

func main() {
	// n is the number of bytes
	// err is a write error
	// both are returned by fmt.Println
	n, err := fmt.Println("Hello everyone, this is the most awesome class in the entire world. We are having fun and learning the GO programming language")

	fmt.Println("Number of bytes:", n)

	fmt.Println("WRITE ERR:", err)

	foo()

	// _ is how we let the compiler know we arent goign to use a return value
	x, _ := fmt.Println("something more")

	fmt.Println("Number of bytes:", x)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
