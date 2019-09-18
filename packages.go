package main

import (
	"fmt"
	"go-tour/computation"
)

//вызов функции из файла соседней директории
func main()  {
	fmt.Println(computation.Factorial(4))
	fmt.Println(computation.Factorial(5))
}
