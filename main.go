package main

import "fmt"

const (
	water = iota
	air
	soil
	glass
)
func main() {
	fmt.Println(water, air, soil, glass)
}
