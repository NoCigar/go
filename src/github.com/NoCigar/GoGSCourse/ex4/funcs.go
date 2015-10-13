package main

/*import(
	"fmt"
)*/

func main() {
	message := "Hello"
	sayHello(&message)
	println(message)
	sayHellos("git", "rekt", "scrub")
}

func sayHello(message *string) {
	println(*message)

	*message = "hello changed inside func"

	println()



}

func sayHellos(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}