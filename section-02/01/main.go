// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	mul := func(a, b int) int {
		return a * b
	}
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(addMulSub(5, 2, add, mul, sub))
	inc := incrementor()
	fmt.Println(inc(), inc(), inc(), inc())
}

func addMulSub(a, b int, add, mul, sub func(a, b int) int) (int, int, int) {
	return add(a, b), mul(a, b), sub(a, b)
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
