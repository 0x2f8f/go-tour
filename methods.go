package main

import "fmt"

type person struct{
	name string
	age int
}

func (p person) print() {
	fmt.Println(p.name, p.age)
}

func main()  {
	var p1 = person{"Serg", 29}
	var p2 = person{"Anastasiya", 35}

	p1.print()
	p2.print()
}
