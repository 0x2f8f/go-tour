package main

import (
	"fmt"
	"sort"
)

func main()  {
	//maps1()
	//mapsLiterals1()
	maps2()
}

type Vertex1 struct {
	Lt, Ln float64
}

func maps1()  {
	var m map[string]Vertex1
	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

//foreach по map
func maps2()  {
	m := make(map[int]string)
	m[2] = "Yakimov"
	m[0] = "Sergey"
	m[1] = "Sergeevich"

	for key, value := range m {
		fmt.Println(key, " ", value)
	}

	//сортируем
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys) //отсортировали
	fmt.Println(keys)
	for _,k := range keys {
		fmt.Println(k, " ", m[k])
	}
}

func mapsLiterals1()  {
	var m = map[string]Vertex1{
		"Bell Labs": Vertex1{
			40.68433, -74.39967,
		},
		"Google": Vertex1{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}