package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var sentence string = "asSASA ddd dsjkdsjs dk"

	onePrintAs()
	fmt.Println()



	twoCountCharsAndOutputBytes(sentence)
	fmt.Println("\n3:")
	fmt.Println("\n "+sentence)

	threeReplaceSubstring(sentence)
	fmt.Println()

	fourReverseString(sentence)
}


func onePrintAs() {

	for i:=0;i<100;i++ {
		for j:=0;j<i;j++{
			fmt.Print("A")
		}
		println()
	}

}

func twoCountCharsAndOutputBytes(sentence string){
	fmt.Printf("\nlength of \"%s\": %d", sentence, len(sentence))
	fmt.Printf("\nlength in bytes %d",utf8.RuneCountInString(sentence))
}

func threeReplaceSubstring(sentence string) {
	fmt.Print("============\nreplaced: "+sentence[0:4]+"-REPLACED-"+
	sentence[7:len(sentence)]+"\n============")
}

func fourReverseString(sentence string) {

	inputByteArr := []byte(sentence)
	var arrSize int = len(inputByteArr)
	outputByteArr := make([]byte,arrSize)
	j:=arrSize

	for i:=0; i<arrSize; i++ {
		outputByteArr[j-i-1] = inputByteArr[i]
	}
	fmt.Println(string(outputByteArr))
}

