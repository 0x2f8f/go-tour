package main

func main() {
	pointers1()
}

func pointers1()  {
	var p *int;
	i := 42
	p = &i
	println(*p)
	i++
	println(*p)
}
