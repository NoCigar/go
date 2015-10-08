package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("first string "+" second string\n")
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("first arg: "+os.Args[0])
	fmt.Println("second arg: ", os.Args[1])
}