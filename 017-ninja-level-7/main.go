package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	//(*p).first = "John"
	//(*p).last = "Wick"
	//(*p).age = 45
	p.first = "John"
	p.last = "Wick"
	p.age = 45
}

func (p person) speak() {
	fmt.Println("My name is", p.last, ",", p.first, p.last, ", and i am", p.age)
}

func main() {

	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p.speak()

	changeMe(&p)
	p.speak()
}
