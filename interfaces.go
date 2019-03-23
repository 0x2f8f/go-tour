package main

import "fmt"

type Vehicle interface{
	move()
}

type Fuel interface {
	refuel()
}

// структура "Автомобиль"
type Car struct{ }

// структура "Самолет"
type Aircraft struct{}

func (c Car) move(){
	fmt.Println("Автомобиль едет")
}

func (c Car) refuel()  {
	fmt.Println("Автомобиль заправляется бензином")
}

func (a Aircraft) move(){
	fmt.Println("Самолет летит")
}

func drive(v Vehicle)  {
	v.move()
}

func main()  {
	//interface1() //структуры Car и Aircraft имплементировали интерфейс Vehicle и реализовали метод move
	interface2() //вызываем отдельную функцию, которая работает с объектами интерфейса
}

//структуры Car и Aircraft имплементировали интерфейс Vehicle и реализовали метод move
func interface1()  {
	var c Vehicle = Car{}
	c.move()

	var a Vehicle = Aircraft{}
	a.move()
}

//вызываем отдельную функцию, которая работает с объектами интерфейса
func interface2() {
	var c Vehicle = Car{}
	var a Vehicle = Aircraft{}
	drive(c)
	drive(a)
}