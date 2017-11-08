package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	for i:=0; i <10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println("solution is %v", z)
	}
    return z
}

func main() {
	fmt.Println("yangu %v", Sqrt(100))
	fmt.Println("yao %v", math.Sqrt(100))
}


