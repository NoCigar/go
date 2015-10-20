package main


func main() {

	print(calcAvg([]float64{1.,2.,3.,4.}))

}


func calcAvg(slice []float64) (result float64){

	var sum float64

	for i:=0;i<len(slice);i++ {
		sum += slice[i]
	}

	return sum/float64(len(slice))
}