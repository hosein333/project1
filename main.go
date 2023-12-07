package main

import "fmt"

func main() {
	i := 52
	var p *int
	p = &i

	fmt.Println("i :", i)
	fmt.Println("p : ", p)
	fmt.Println("&i : ", &i)
	fmt.Println("*p : ", *p)

	i = 102

	fmt.Println("*p :", *p)



}
