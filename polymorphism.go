package main
import "fmt"

type Vehicle2 interface{
	move()
}

type Car2 struct{ model string}
type Aircraft2 struct{ model string}


func (c Car2) move(){
	fmt.Println(c.model, "едет")
}
func (a Aircraft2) move(){
	fmt.Println(a.model, "летит")
}

/**
  Полиморфизм
  В зависимости от реального типа структуры динамически определяется,
  какая именно реализация метода move для какой структуры должна вызываться.
 */
func main() {

	tesla := Car2{"Tesla"}
	volvo := Car2{"Volvo"}
	boeing := Aircraft2{"Boeing"}

	vehicles := [...]Vehicle2{tesla, volvo, boeing}
	for _, vehicle := range vehicles{
		vehicle.move()
	}
}