package main

import "fmt"

func main() {
	// arrays
	var a [5]int
	fmt.Println("empty: ", a)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println("full: ", a)

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("type: %T\n", b)
	for i, v := range b {
		fmt.Printf("location is %v, value is %v\n", b[i], v)
	}

	// slices
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Printf("location of x is %v, value is %v\n", i, v)
	}

	// partial indexing
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	// append
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	j := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	j = append(j[:3], j[6:]...)
	fmt.Println(j)

	// slice using make
	m := make([]string, 50, 50)
	m = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	for i := 0; i < len(m); i++ {
		fmt.Println(i, m[i])
	}

	// multi dimensional slices
	k := []string{
		"James",
		"Bond",
		"Shaken, not stirred",
	}
	p := []string{
		"Miss",
		"Moneypenny",
		"Helloooooo, James",
	}
	o := [][]string{k, p}
	for i, v := range o {
		fmt.Println(i, v)
		for g, h := range v {
			fmt.Println(g, h)
		}
	}

	// maps
	mp := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Beimg evil", "Ice cream", "Sunsets"},
	}
	mp["smith_greg"] = []string{"Making money", "Puppies", "Sunshine"}

	delete(mp, "moneypenny_miss")

	for i, j := range mp {
		fmt.Println("This is the record for: ", i)

		for k, v2 := range j {
			fmt.Println("\t", k, v2)
		}
	}

}
