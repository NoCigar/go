package main
import "fmt"


func main() {

	list := []int{5,-1,12,3,5,0}
	sortedList:=bubbleSort(list)
	fmt.Println(list)
	fmt.Println("\n",sortedList)
}

func bubbleSort(input []int) []int {

	var swapped bool = true



	for swapped == true {
		swapped = false
		for i:=0;i<len(input)-1;i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}

	return input
}