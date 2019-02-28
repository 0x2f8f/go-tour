package main

import "fmt"

func main()  {
	vertex1();
	vertex2();
	vertex3();
}

//литералы структур
func vertex3()  {
	v1 := Vertex{1,2}
	v2 := Vertex{x: 1}
	v3 := Vertex{}
	p := &Vertex{3,4}
	p.y=5

	fmt.Println(v1, v2, v3, p)
}

//структуры
func vertex1()  {
	var v = Vertex{1,2}
	fmt.Println(v)
	fmt.Println(v.y)

	v.y = 3
	fmt.Println(v.y)

	v.y++
	fmt.Println(v.y)
}

//указатели на структуры
func vertex2()  {
	v := Vertex{1,2}
	p := &v
	p.x = 3
	(*p).y=4
	fmt.Println(v)
}

type Vertex struct {
	x int;
	y int;
}