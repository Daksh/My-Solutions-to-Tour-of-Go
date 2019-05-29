package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x<0{
		return 0, ErrNegativeSqrt(x)		
	}
	ans := x/2
	iter := 0 
	for ep := 0.000000000001; ans*ans - x > ep || ans*ans-x < -1*ep; {
		ans -= (ans*ans - x) / (2*ans)
		iter++
	}
	//fmt.Println("Iterations:",iter)
	return ans, nil

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
