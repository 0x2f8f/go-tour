package main

import "fmt"

func main() {
	pointers1()
	pointers2()
}

//указатели
func pointers1()  {
	var p *int;
	i := 42
	p = &i
	println(*p) //42

	i++
	println(*p) //43

	*p = 40
	println(i) //40
	i++
	println(*p) //41
}

//передача по ссылке
func pointers2() {
	x := 5
	zero2(&x)
	fmt.Println(x) // x is 0
}

func zero2(xPtr *int) {
	*xPtr = 0
}