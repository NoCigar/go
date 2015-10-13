package main
import "fmt"

func main() {
	foo := myStruct{}
	foo.myField = "asd"
	println(foo.myField)

	bar := myStruct{"qwe"}
	println(bar.myField)

	ter := new(myStruct)
	ter.myField="zxc"
	println(ter.myField)

	ter2 := &myStruct{"cxz"}

	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(ter)
	println(ter)
	fmt.Println(*ter)
	println()
	fmt.Println(ter2)
	println(ter2)
	fmt.Println(*ter2)

}

type myStruct struct {
	myField string
}