package main

import "fmt"
type float float64

func main() {
	var f float = 5.2
	var g float64 = 5.2

	fmt.Println("f == g", f == float(g))
}
