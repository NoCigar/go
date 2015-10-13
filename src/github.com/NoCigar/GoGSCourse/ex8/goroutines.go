package main

import (
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	go abcGen()

	println("main goes first!")

	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for I := byte('a'); I <= byte('z'); I++ {
		println(string(I))
	}
}