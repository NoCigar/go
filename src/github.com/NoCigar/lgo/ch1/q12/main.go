package main
import (
	"fmt"
	"strings"
)

func main() {

	myList := []int{1,2,3,4,5}
	returnedList := []int{}
	f := func (input int) int {
		return input*2
	}
	returnedList = mapStuff(f, myList)
	fmt.Print(returnedList)

	myStringList := []string{"aaa", "bbb", "ccc"}
	returnedStringList := []string{}

	f2 := func (input string) string {
		return strings.ToUpper(input)
	}

	returnedStringList = mapStuffString(f2, myStringList)
	fmt.Print("\n",returnedStringList)
}


func mapStuff( modify func(int) int, list []int) []int {

	n := len(list)
	modifiedList := make([]int,n)
	for i:=0; i<n; i++ {
		modifiedList[i] = modify(list[i])
	}

	return modifiedList
}

func mapStuffString ( modify func(string) string, list []string) []string {
	n := len(list)
	modifiedList := make([]string,n)

	for i:=0; i<n; i++ {
		modifiedList[i] = modify(list[i])
	}
	return modifiedList
}

