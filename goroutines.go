package main

import (
	"fmt"
	"time"
)

func main()  {
	gor1()
}

var counter = 0

func gor1()  {
	for i := 0; i < 10; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

//
func incr() {
	counter++
	fmt.Println(counter)
}