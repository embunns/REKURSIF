package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(gcd(a, b))
}
func gcd(a, b int)int{
	var c int
	if a%b==0{
		return b
	}else{
		c=a%b
		return gcd(b, c)
	}
}
