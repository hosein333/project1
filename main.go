package main

import "fmt"

func main() {
	i := 52
	var p *int
	p = &i

	fmt.Println(i)
	fmt.Println(p)
}
