package main
import "fmt"

func main() {
	printInts(1,2,3,4,5,12)
}

func printInts(args ...int) {

	for i:= range args {
		fmt.Println(args[i])
	}
}