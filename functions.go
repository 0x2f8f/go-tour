package main

import "fmt"

func main()  {
	functions1()
}

func functions1()  {
	fmt.Println(add(42, 13))
}

//складывание чисел
func add(x int, y int) int {
	return x + y
}