package main

import "fmt"

func main() {

	var (
		var1 float64 = -1.2
		var2 int = 1
	)
	var3 := var1 + float64(var2)
	
	if var3 == -0.2 {
		fmt.Println("It is Equal")
	}

	fmt.Println(var1, var2, var3,)
}

