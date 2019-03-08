package main

import (
	"fmt"
)

func main()  {
	//channels1()
	//channels2()
	//channels3()
	//channels4()
	//channels5()
	//channels6()
	channels7()
}

//канал с типом int
func channels1()  {
	var intCh chan int = make(chan int) // канал для данных типа int

	go func(){
		intCh <- 5 // блокировка, пока данные не будут получены функцией main
	}()
	a := <- intCh
	fmt.Println(a)
}

//канал с типом string
func channels2()  {
	strCh := make(chan string)  // канал для данных типа string

	go func(){
		strCh <- "Serg Yakimov" // блокировка, пока данные не будут получены функцией main
	}()
	s := <- strCh
	fmt.Println(s)
}

//смотрим на почледовательность выполнения анонимныз горутин
func channels3()  {
	strCh := make(chan string)  // канал для данных типа string

	fmt.Println(1)
	go func(){
		fmt.Println(3)
		strCh <- "Serg Yakimov" // блокировка, пока данные не будут получены функцией main
		fmt.Println(4)
	}()
	fmt.Println(2)
	s := <- strCh
	fmt.Println(5)
	fmt.Println(s)
	fmt.Println(6)
}

//передача канала в горутину
func channels4()  {
	intCh := make(chan int)

	go factorial(5, intCh)  // вызов горутины
	fmt.Println(<-intCh) // получение данных из канала
	fmt.Println("The End")
}

//функция факториала
func factorial(n int, ch chan int){
	result := 1
	for i := 1; i <= n; i++{
		result *= i
	}
	fmt.Println(n, "-", result)

	ch <- result     // отправка данных в канал
}

//Буферизированные каналы
func channels5()  {
	intCh := make(chan int, 3)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	fmt.Println(<-intCh)    // 1
	fmt.Println(<-intCh)    // 2
	fmt.Println(<-intCh)	// 3

	intCh <- 4
	a := <- intCh
	fmt.Println(a)			// 4
}

//блокировка при переполнении буфера (будещ ошибка)
func channels6()  {
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	intCh <- 24

	//дальше будет ошибка
	intCh <- 15  // блокировка - функция main ждет, когда освободится место в канале

	fmt.Println(<-intCh)
	fmt.Println("The End")
}

func channels7()  {
	intCh := make(chan int, 3)
	intCh <- 10

	//емкость
	fmt.Println(cap(intCh))     // 3

	//кол-во элементов в канале
	fmt.Println(len(intCh))     // 1

	fmt.Println(<-intCh)
}