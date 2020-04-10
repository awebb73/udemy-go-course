package main

import "fmt"

const (
	a = 1 << iota // a == 1 (iota == 0)
	b = 1 << iota // b == 2 (iota == 1)
	c = 3         // c == 3 (iota == 2)
	d = 1 << iota // d == 8 (iota == 3)
)

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0 	(iota == 0)
	bit1, mask1                          // bit1 == 2, mask0 == 1 	(iota == 1)
	_, _                                 // 						(iota == 2, unused)
	bit3, mask3                          // bit3 == 8, mask3 == 7 	(iota == 3)
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println("")

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println("")

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
	fmt.Println("")

	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#x\n", i, v)
	}
	fmt.Println("")

	fmt.Printf("%d\t%b\n", a, a)
	fmt.Printf("%d\t%b\n", b, b)
	fmt.Printf("%d\t%b\n", c, c)
	fmt.Printf("%d\t%b\n", d, d)
	fmt.Println("")

	fmt.Printf("%d\t%b\t%d\t%b\n", bit0, bit0, mask0, mask0)
	fmt.Printf("%d\t%b\t%d\t%b\n", bit1, bit1, mask1, mask1)
	fmt.Printf("%d\t%b\t%d\t%b\n", bit3, bit3, mask3, mask3)
	fmt.Println("")

}
