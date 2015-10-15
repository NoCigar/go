package main
import "fmt"


func main() {

	fmt.Println("\nfor loop print:")

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nGo to print:")

	i:=0
	Here:
		fmt.Println(i)
		if i<10 {
			i++
			goto Here
		}

	fmt.Println("\nfor loop array print")

	var arr [10]int

	for i:=0; i < 10; i++ {
		arr[i] = i
		fmt.Println(arr[i])
	}
	fmt.Println(arr)
}