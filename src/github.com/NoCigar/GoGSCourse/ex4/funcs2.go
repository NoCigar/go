package main

func main(){

	result := add(1,3,5,8)
	println(result)

	numTerms,sum := addWith2Results(1,3,5,8)

	println("Added: ",numTerms, "items and their sum is: ", sum)

	numTerms,sum = addWith2ResultsNames(1,3,5,10)

	println("Added: ",numTerms, "items and their sum is: ", sum)

}

func add(terms ...int) int {

	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

func addWith2Results(terms ...int) (int, int) {
	result:= 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

func addWith2ResultsNames(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}

	numTerms = len(terms)

	return
}

