package main

import "fmt"

func main()  {
	functions1()
}

func functions1()  {
	fmt.Println(add2(42, 13))
}

//складывание чисел
func add2(x int, y int) int {
	return x + y
}