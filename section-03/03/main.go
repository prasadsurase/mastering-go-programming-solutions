package main

import "fmt"

func main() {
	ch := make(chan string)
	go Sender(ch)
	for str := range ch {
		fmt.Println("Received:", str)
	}
	if _, ok := <-ch; !ok {
		fmt.Println("Channel Closed")
	} else {
		fmt.Println("Channel Open")
	}
}

func Sender(ch chan string) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- "hello"
	}
}
