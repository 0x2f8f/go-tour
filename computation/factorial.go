package computation

//если объявить с маленькой, то функция в других файлах будет не доступна
func Factorial(n int) int {
	var result = 1
	for i:=1; i <= n; i++{
		result *= i
	}
	return result
}