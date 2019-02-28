package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	for1()
	for2()
	for3()
	while1()
	if1()
	if2()
	switch1()
	switch2()
	switch3()
	switch4()
}

func switch4() {
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		println("доброе утро");
	case hour <17:
		println("добрый день")
	default:
		println("добрый вечер")
	}
}

func switch3() {
	today := time.Now().Weekday()
	monday := time.Monday;
	switch monday {
	case today + 0:
		println("Monday today")
	case today + 1:
		println("Monday Tomorrow")
	case today + 2:
		println("Monday In two days")
	default:
		println("Monday Too far away")
	}
}

func switch2() {
	//os := "ubuntu"
	switch os:= runtime.GOOS; os {
	case "linux":
		println("linux")
	case "XP":
		println("XP")
	default:
		println(os)
	}
}

func switch1() {
	os := "ubuntu"
	switch os {
		case "linux":
			println("linux")
		case "XP":
			println("XP")
		default:
			println(os)
	}
}

func if2() {
	i :=1
	if j:=-i; j<0 {
		println("меньше нуля");
	} else {
		println("не меньше нуля");
	}
}

func for1() {
	sum := 0
	for i:=0; i<10; i++ {
		sum+=i
	}
	println(sum)
}

func for2() {
	sum := 1
	for ; sum<100; {
		sum+=sum
	}
	println(sum)
}

func for3() {
	sum := 1
	for ; sum<100;sum+=sum {}
	println(sum)
}

func while1() {
	sum :=3
	for sum<100 {
		sum += sum
	}
	println(sum)
}

func if1() {
	println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}