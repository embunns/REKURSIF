package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	ganjil(x)

}

func ganjil(a int) {
	if a == 1 {
		fmt.Print(a)
	} else {
		if a%2 == 0 {
			ganjil(a - 1)
		} else {
			ganjil(a - 2)
			fmt.Print(a)
		}
	}
}
