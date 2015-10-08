package main

import(
	"fmt"
)

func main(){

	noResult()
	fmt.Printf("\nsingleResult returns: %d", singleResult(5,3))

	_, boolval := multipleResults("first")

	fmt.Printf("\nmultiple result first test:  %t", boolval)


	value, boolval := multipleResults("second")
	fmt.Printf("\nmultiple result second test: %d, %t", value, boolval )

}

func noResult(){
	fmt.Printf("\nthis function returns nothing")
}

func singleResult(a int, b int) int {
	return a+b
}

func multipleResults(name string) (int, bool){
	if name == "first"{
		return 42, true
	} else {
		return 24, false
	}
}