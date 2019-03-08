package main

import "fmt"

func main()  {
	//array1()
	//array2()
	//array3()
}

func array1()  {
	var a [10]int;
	fmt.Println(a);
}

func array2()  {
	var a [10]int;
	a[1] = 1
	a[2] = 2
	fmt.Println(a);
}

func array3()  {
	a := [4]int{1,2,3}
	fmt.Println(a);
}