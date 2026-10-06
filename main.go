package main

import "fmt"

func main() {
	i := 15

	p := &i

	fmt.Println(*p)

	*p = 300

	fmt.Println(i)
}
