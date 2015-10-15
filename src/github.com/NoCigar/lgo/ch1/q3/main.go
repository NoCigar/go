package main
import "fmt"


func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i:=1;i<101;i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}
}
