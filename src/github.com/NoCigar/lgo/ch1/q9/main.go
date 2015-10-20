package main
import (
	"fmt"
	"strconv"
)


func main() {
	myStack := new(stack)
	myStack.push(1)
	myStack.push(2)
	myStack.push(3)
	fmt.Printf("stack %v\n",myStack)
	myStack.pop()
	myStack.pop()
	fmt.Printf("stack %v\n",myStack)
}


type stack struct {
	index int
	slice [10]int
}

func (s *stack) push(element int){
	if(s.index>len(s.slice)){
		return
	}
 	s.slice[s.index] = element
	s.index++

}

func (s *stack) pop() (res int){
	if s.index == 0{
		return
	}
	s.index--
	return s.slice[s.index]
}

func (s stack) String() string {
	var res string
	for i:=0; i<s.index; i++ {
		res = "( "+ strconv.Itoa(i)+" , "+
			strconv.Itoa(s.slice[i]) + " )"
	}
	return res
}