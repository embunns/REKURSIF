package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(0, 1, " ")
	fibonacci(0, 1, n)
}
func fibonacci(a, b, n int) {
	var sum int
	if n >= 2 {
		sum = a + b
		fmt.Print(sum, " ")
		a = b
		b = sum
		fibonacci(a, b, n-1)
	}
}
