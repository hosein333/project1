package main

import (
	"fmt"
        "math"
)

const float64EqualityThreshold = 1e-9

func main() {



	var (
		var1 float64 = -1.2
		var2 int = 1
	)
	var3 := var1 + float64(var2)
	
	if math.Abs(var3-(-0.2)) < float64EqualityThreshold {
		fmt.Println("It is Equal")
	}

	fmt.Println(var1, var2, var3,)
}

