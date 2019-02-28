package main

import (
	"fmt"
)

func main() {
//	fmt.Println("Hello, world")
//	fmt.Println("Current time:", time.Now())
//	fmt.Println("Random number:", rand.Intn(10))
//	fmt.Println(math.Pi)
//	fmt.Println(add(2,3))
//	testSwap()
	fmt.Println(split(17))
}

func split(sum int) (y, z int) {
	var x = sum * 4 / 9
	y = sum - x
	z = x+y
	return
}

//метод добавления двух чисел
func add(x, y int) int {
	return x+y;
}

//тестим получение слов из другой функции
func testSwap() {
	a, b := swap("hello", "world")
	fmt.Println(a, b);
}

//возвращает 2 слова
func swap(x, y string) (string, string) {
	return x,y;
}

/*
Run
$ go run hello.go

Compile
$ go install hello.go

Run app
$ hello
*/
