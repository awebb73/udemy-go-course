package main

import (
	"fmt"
)

const (
	g int = 73
	h     = "Hello, World"
)

const (
	yr1 = 2020 + iota
	yr2
	yr3
	yr4
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	fmt.Println("")

	a := 3 == 3
	b := 3 <= 2
	c := 4 >= 3
	d := 5 != 2
	e := 3 < 4
	f := 4 > 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("")

	fmt.Println(g)
	fmt.Printf("%T\n", g)
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	fmt.Println("")

	y := 42
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
	w := y << 1
	fmt.Printf("%d\t%b\t%#x\n", w, w, w)
	fmt.Println("")

	str1 := `this is a raw "string" literal`
	fmt.Println(str1)
	fmt.Println("")

	fmt.Println(yr1)
	fmt.Println(yr2)
	fmt.Println(yr3)
	fmt.Println(yr4)
	fmt.Println("")
}
