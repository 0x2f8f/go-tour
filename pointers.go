package main

func main() {
	pointers1()
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
