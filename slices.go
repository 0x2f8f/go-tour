package main

import "fmt"

func main()  {
	slices1();
	slices2();
	slices3();
	slices4();
	slices5();
	slices6();
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

func slices6() {
	println("------")
	s := []int{0,1,2,3,4,5}
	printSlice(s)

	s = s[:3]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
