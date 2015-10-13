package main
import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	ch := make(chan string)
	doneCh := make(chan bool)

	go abcGen2(ch)
	go printer(ch,doneCh)

	<-doneCh

	//time.Sleep(100* time.Millisecond)
}

func abcGen2(ch chan string) {
	for I := byte('a');  I<= byte('z'); I++ {
		ch <- string(I)
	}

	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for I:= range ch{
		println(I)
	}
	doneCh <- true
}
