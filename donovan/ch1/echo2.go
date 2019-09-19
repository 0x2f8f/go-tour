package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command name: " +os.Args[0])
	for ind, arg := range os.Args[1:] {
		fmt.Println(ind,": ",arg)
	}
}