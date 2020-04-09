package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world. We are having fun and learning the GO programming language")

	foo()

	fmt.Println("something more")

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
