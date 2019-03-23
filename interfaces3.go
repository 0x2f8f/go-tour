package main

import "fmt"

type Reader interface{
	read()
}

type Writer interface{
	write(string)
}

func writeToStream2(writer Writer, text string){
	writer.write(text)
}
func readFromStream2(reader Reader){
	reader.read()
}

type File2 struct{
	text string
}

func (f *File2) read(){
	fmt.Println(f.text)
}
func (f *File2) write(message string){
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

/**
  Реализация нескольких интерфейсов
 */
func main() {

	myFile := &File2{}
	writeToStream2(myFile, "hello world")
	readFromStream2(myFile)
}