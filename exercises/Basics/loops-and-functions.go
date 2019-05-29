package main

import (
	"fmt"; "math"
)

func Sqrt(x float64) float64 {
	ans := x/2
	iter := 0 
	for ep := 0.000000000001; ans*ans - x > ep || ans*ans-x < -1*ep; {
		ans -= (ans*ans - x) / (2*ans)
		iter++
	}
	fmt.Println("Iterations:",iter)
	return ans
}

func main() {
	var num = 2232.0
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}
