package main

import "fmt"

func main()  {
	maps1()
	mapsLiterals1()
}

type Vertex struct {
	Lt, Ln float64
}

func maps1()  {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*
func maps2()  {
	m := map[int]string
	m[0]="Serg"
}
*/

func mapsLiterals1()  {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}