package main

import "fmt"

type person struct {
	first string
	last  string
	cream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		cream: []string{"chocolate", "strawberry", "vanilla"},
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		cream: []string{"rockyroad", "moosetracks", "cherry garcia"},
	}

	g := []person{p1, p2}

	for i, j := range g {
		fmt.Printf("index is: %v,\tvalue is: %v\n", i, j)

	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, j := range p1.cream {
		fmt.Println(i, j)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, j := range p2.cream {
		fmt.Println(i, j)
	}

	// with maps
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, j := range m {
		fmt.Println("Name: ", i)
		for k, l := range j.cream {
			fmt.Printf("\tindex: %v,\tvalue: %v\n", k, l)
		}
	}

	// imbedded strucs
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: false,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(t.vehicle.doors)
	fmt.Println(t.vehicle.color)
	fmt.Println(t.fourWheel)
	fmt.Println(s)
	fmt.Println(s.vehicle.doors)
	fmt.Println(s.vehicle.color)
	fmt.Println(s.luxury)

	z := struct {
		first   string
		friends map[string]int
		drinks  []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 001,
			"Q":          002,
			"M":          003,
		},
		drinks: []string{"Martini", "Whiskey", "Gin", "Vodka"},
	}
	fmt.Println(z)
}
