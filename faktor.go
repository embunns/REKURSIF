package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	faktor(1, n)
}
func faktor(x, n int) {
	if n == x {
		fmt.Print(n)
	} else {
		if n%x == 0 {
			fmt.Print(x, " ")
		}
		faktor(x+1, n)
	}
}
