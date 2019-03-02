package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main()  {
	slices1();
	slices2();
	slices3();
	slices4();
	slices5();
	slices6();
	slices7();

	nil1();
	make1();
	append1();

	range1();
	range2();
	println("------")
	range3();
}

func slices1()  {
	var a [10]int; 	//массив
	var s []int;	//слайс (он же срез)
	fmt.Println(a, s);
}

//несколько способов объявления слайсов
//смотрим типы
func slices2()  {
	a := [6]int{0,1,2,3,4,5}
	var s []int = a[0:3];
	s2 := a[3:6]

	fmt.Println(a, s, s2);
	fmt.Printf("%T %T %T\n", a, s, s2)
}

//слайсы - это указатели на массивы
//при изменении элемента массива, меняется значение и у слайса
func slices3()  {
	a := [6]int{0,1,2,3,4,5}
	var s1 []int = a[0:3]
	var s2 []int = a[2:4]
	fmt.Println(a, s1, s2);

	a[2] = 22
	fmt.Println(a, s1, s2);
}

//литералы срезов
//в этих ситуациях на самом деле вначале в памяти создаётся массив,
// а после слайс, который на негол ссылается
func slices4()  {
	si := []int{1, 2, 3}
	sb := []bool{true, true, false}
	ss := []string{"one", "two", "three"}
	fmt.Println(si, sb, ss)

	st := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(st)
}

//без указания верхней или нижней границы
func slices5() {
	a := [6]int{0,1,2,3,4,5}
	s1 := a[1:]
	s2 := a[:3]
	fmt.Println(s1, s2)
}

//размеры и вместимость в слайсах
func slices6() {
	s := []int{0,1,2,3,4,5}
	printSlice(s)

	s = s[:3]
	printSlice(s)

	var d = s[:1]
	printSlice(d)
}

//двумерный слайс. Слайс состоит из слайсов
func slices7()  {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][2] = "X"
	board[1][0] = "O"
	board[1][2] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//нулевой слайс
func nil1() {
	var s []int
	printSlice(s)
	if (s == nil) {
		fmt.Printf("nil! %T\n", s)
	}
}

//создание через make
func make1()  {
	a := make([]int, 5)  // len(a)=5, cap(a)=5
	printSlice(a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice(b)

	c := b[:2] // len(c)=2, cap(c)=5
	printSlice(c)

	d := c[2:5] // len(d)=3, len(d)=3
	printSlice(d)

	e := make([]int, 4, 5) // len(e)=4, cap(e)=5
	printSlice(e)
}

//добавлнение элемента в слайс
func append1()  {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

//цицкл for с range, к-ый возвращает индекс и копию элемента
//работает как foreach
func range1()  {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//range без индекса
func range2()  {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func range3()  {
	sl := Pic(4,4)
	for i, v := range sl {
		fmt.Printf("%d, %d\n", i, v)
	}
}

//dy - размер слайса
//dx - размер внутреннних слайсов
//внутренний слайс состоит из 8-битных беззнаковых целых чисел
//func Pic(dx, dy int) [][]uint8 {
func Pic(dx, dy int) [][]uint8 {
	sl := [][]uint8{}

	for i:=0;i<dx;i++ {
		kl := []uint8{}
		for j:=0;j<dy;j++ {
			kl = append(kl, uint8(rand.Intn(10)))
		}
		sl = append(sl, kl)
	}

	return sl
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}