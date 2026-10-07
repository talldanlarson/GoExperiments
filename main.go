package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
}

func main() {
	m["Google"] = Vertex{
		37.42202, -112.08408,
	}
	fmt.Println(m)

	delete(m, "Google")
	fmt.Println(m)

	v, ok := m["Google"]
	fmt.Println("The value:", v, "Present?", ok)
}
