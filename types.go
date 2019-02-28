package main

import "fmt"

var (
	tobe bool = false
	MaxInteger uint64 = 1<<64 -1
	F float64 = 1<<64 -1
)

const Pi = 3.14

func main () {
	//types()
	typeConversions()
	typeConversions2()
	typeConversions3()
	println(Pi)
}

//приведение типа
func typeConversions() {
	var i int = 42
	var f float32 = float32(i)
	var u uint = uint(f)
	println(i, f, u)
}

//короткая запись приведения типа
func typeConversions2() {
	i := 42
	f := float64(i)
	u := uint(f)
	println(i, f, u)
}

//выведение типа
func typeConversions3() {
	i := 42
	f := 3.142
	h := 0.867 + 0.5i
	s := "3.142"
	fmt.Printf("%T %T %T %T\n", i, f, h, s)
}

//максимальные значения типов
func types() {
	fmt.Printf("Type: %T, value: %v\n", tobe, tobe)
	fmt.Printf("Type: %T, value: %v\n", MaxInteger, MaxInteger)
	fmt.Printf("Type: %T, value: %v\n", F, F)
}

/**
Типы в GO:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // псевдоним для uint8

rune // псевдоним для int32
     // представляет Unicode код

float32 float64

complex64 complex128
 */
