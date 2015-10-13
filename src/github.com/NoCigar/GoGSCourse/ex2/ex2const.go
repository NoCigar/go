package main

const (
	first = iota << (10 * iota)
	second
)

const (
	third = iota
)

func main() {
	println(first)
	println(second)
	println(third)

}
