package main

import "fmt"

type Reader4 interface{
	read()
}

type Writer4 interface{
	write(string)
}

type ReaderWriter interface{
	Reader4
	Writer4
}

func writeToStream4(writer ReaderWriter, text string){
	writer.write(text)
}
func readFromStream4(reader ReaderWriter){
	reader.read()
}

type File4 struct{
	text string
}

func (f *File4) read(){
	fmt.Println(f.text)
}
func (f *File4) write(message string){
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

/**
  Вложенные интерфейсы
 */
func main() {

	myFile := &File4{}
	writeToStream4(myFile, "hello world")
	readFromStream4(myFile)
	writeToStream4(myFile, "lolly bomb")
	readFromStream4(myFile)
}