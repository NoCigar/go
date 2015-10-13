package main

func main(){


	addFunc:=func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}

		numTerms = len(terms)

		return
	}

	items,sum := addFunc(1,2,3,4)

	println("Added: ",items, "items and their sum is: ", sum)



}




