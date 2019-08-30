package main

import (
	"fmt"
)

func reverse(num int)(int){
	new_num := 0
	for num != 0{
		new_num *= 10
		new_num += num%10
		num /= 10
	}
	return new_num
}

func main(){
	x := 1234
	y := reverse(x)
	fmt.Println(y)
}

