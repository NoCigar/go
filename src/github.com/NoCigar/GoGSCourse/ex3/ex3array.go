package main

import "fmt"

func main() {
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 32
	myArray[2] = 99

	mySlice := myArray[:]

	fmt.Println(myArray)
	fmt.Println(len(myArray))
	fmt.Println(mySlice)

	mySlice = append(mySlice, 100)
	fmt.Println(mySlice)

	//declare slice
	mySlice2 := []float32{12., 15., 18.}
	fmt.Println(mySlice2)

	myMadeSlice := make([]float32, 100)
	myMadeSlice[0] = 13.
	myMadeSlice[1] = 15.
	fmt.Println(myMadeSlice)

	//map
	myMap := make(map[int]string)
	fmt.Println(myMap)

	myMap[42] = "fug"
	myMap[12] = "meh"
	fmt.Println(myMap[42])
	fmt.Println(myMap[123])

}
