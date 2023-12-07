package main

import "fmt"

const (
	water = iota
	air = iota
	soil = iota
	glass = iota
)
func main() {
	fmt.Println(water, air, soil, glass)
}
