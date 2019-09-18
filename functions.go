package main

import "fmt"

func main()  {
	functions1()
	functions2()
}

func functions1()  {
	fmt.Println(add2(42, 13))
}

//складывание чисел
func add2(x int, y int) int {
	return x + y
}

//передача по ссылке
func functions2() {
	x := 5
	zero1(&x)
	fmt.Println(x) // x is 0
}

func zero1(xPtr *int) {
	*xPtr = 0
}