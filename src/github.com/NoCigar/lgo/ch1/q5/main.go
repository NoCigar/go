package main
import "fmt"

func main() {
	arr  := []float64{2.5,2.5,5,10}
	fmt.Printf("Reuslt = %f", calcAvg(arr))
}

func calcAvg(slice []float64) float64 {

	var avg float64 = 0.
	for i:=0; i<len(slice); i++ {
		avg += slice[i]
	}

	return avg/float64(len(slice))
}