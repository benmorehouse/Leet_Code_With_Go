package main

import "fmt"

func myPow(x float64, n int) float64 {
	// the most bruitforce way to do this is looping through and iterating
	// but we can do better and split it in halves
	neg := false
	if n==0{
		return 1
	}else if n<0{
		n*=-1
		neg = true
	}
	expo := 1
	base := x
	for expo <= n/2{
		base *= base
		expo *= 2
		fmt.Println(expo)
	}

	fmt.Println(n-expo%n)
	if expo != n{
		for i:=0;i<n-expo%n;i++{
			base*=x
		}
	}

	if neg{
		return 1/base
	}

	return base
}

func main(){
	fmt.Println(myPow(2,-1))
}



