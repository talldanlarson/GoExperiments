package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	b[0] = "XXX"
	fmt.Println(b, c)
	fmt.Println(names)

	t := []int{2, 3, 5, 7, 11, 13}

	t = t[1:4]
	fmt.Println(t)

	t = t[:2]
	fmt.Println(t)

	t = t[1:]
	fmt.Println(t)

	printSlice(s)

	var u []int
	printSlice(u)
	if u == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
