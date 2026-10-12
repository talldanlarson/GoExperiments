package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	var t *T

	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var g interface{}
	describe2(g)
	g = 42
	describe2(g)
	g = "hello"
	describe2(g)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(x interface{}) {
	fmt.Printf("(%v, %T)\n", x, x)
}
