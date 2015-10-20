package main
import "fmt"

func main() {
	fmt.Println(getOrderedInts(2,5))
	fmt.Println(getOrderedInts(6,1))
	fmt.Println(getOrderedInts(1,1))
}

func getOrderedInts(a int, b int) (res1 int, res2 int) {

	if a > b {
		return b,a
	} else {
		return a,b
	}

}