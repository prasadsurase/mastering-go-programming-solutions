package main

import "fmt"

func main() {
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("hello")
	c <- true
	<-c
}

func waitAndSay(c chan bool, message string) {
	if b := <-c; b {
		fmt.Println(message)
	}
	c <- true
}
