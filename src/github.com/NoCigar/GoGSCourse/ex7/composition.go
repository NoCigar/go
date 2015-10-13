package main


func main() {
	mp := messagePrinter{"foo"}
	mp.printMessage()

	mp2 := enhancedMessagePrinter{messagePrinter{"foo2"}}
	mp2.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}