package main


func main() {

	mySlice := make([]int,5)
	mySlice = append(mySlice,1,0,5,9,123)
	println(min(mySlice))
	println(max(mySlice))
}

func min(input []int) int {
	min := input[0]
	for i:=1; i<len(input); i++ {
		if input[i]<min {
			min=input[i]
		}
	}
	return min
}

func max(input []int) int {
	max := input[0]
	for i:=1; i<len(input); i++ {
		if input[i]>max {
			max = input[i]
		}
	}
	return max
}