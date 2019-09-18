package main

import "fmt"

type person struct{
	name string
	age int
}

func (p person) print() {
	fmt.Println(p.name, p.age)
}

func (p *person) updateAge(newAge int){
	(*p).age = newAge
}

func (p person) updateAge2(newAge int){
	p.age = newAge
}

func main()  {
	var p1 = person{"Serg", 29}
	var p2 = person{"Anastasiya", 35}

	p1.print()
	p2.print()

	p2.updateAge(36) //поменяется на 36
	p2.print()

	p2.updateAge2(37) //ничего не произойдет, останется 36, т.к. передастся копия объекта
	p2.print()

	//указатель на объект
	var p3 *person = &p2

	p3.updateAge(37) //поменятся на 37, т.к передан указатель
	p2.print()
}
