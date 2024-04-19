package main

import (
	"fmt"
	"math/rand"
	"time"
)

var points = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
}

func getPoint(name string, c chan int) {
	time.Sleep(time.Duration(time.Duration(rand.Intn(30)) * time.Second))
	c <- points[name]
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go getPoint("A", ch1)
	go getPoint("E", ch2)

	select {
	case p := <-ch1:
		fmt.Printf("%v has %d points\n", "A", p)
	case p := <-ch2:
		fmt.Printf("%v has %d points\n", "E", p)
	case <-time.After(5 * time.Second):
		fmt.Println("Search timeout!")
	}
}
