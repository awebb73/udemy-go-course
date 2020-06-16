package main

import "fmt"

func main() {

	for i := 1; i <= 10000; i++ {
		fmt.Println("The value of i: ", i)
	}

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	for bd := 1980; bd <= 2020; {
		fmt.Println(bd)
		bd++
	}
	bd := 1980
	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}

	for i := 10; i < 101; i++ {
		fmt.Println("mod: ", i%4)
	}
	y := "yellow"
	if y == "blue" {
		fmt.Println("blue")
	} else if y == "green" {
		fmt.Println("green")
	} else {
		fmt.Println("neither")
	}

	switch {
	case false:
		fmt.Println("False")
	case true:
		fmt.Println("True")
	}

	favSport := "soccer"
	switch favSport {
	case "soccer":
		fmt.Printf("favSport is: %v\n", favSport)
	case "football":
		fmt.Printf("favSport is(football): %v\n ", favSport)
	}
}
