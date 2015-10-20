package main
import "fmt"


func main() {

	fmt.Printf("%v",fib(4))

}


func fib(n int) ([]int){
	var i int
	res := make([]int, n)
	res[0] = 1
	if n<2 {
		return res
	}
	res[1] = 1
	for i=2;i<n;i++ {
		res[i]=res[i-1]+res[i-2]
	}

	return res
}