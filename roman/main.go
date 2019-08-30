package main

import(
	"fmt"
)

var rmn = map[byte]int{
	'I' : 1,
	'V' : 5,
	'X' : 10,
}

func roman(s string)(int){
	temp:=[]byte(s)
	new_num := 0
	for i:=len(s)-1;i>-1;i--{
		if rmn[temp[i-1]] > rmn[temp[i]] && i-1 >= 0{
			if rmn[temp[i-1]]>rmn[temp[i]]{
				new_num+=rmn[temp[i-1]] - rmn[temp[i]]
			}
		}else{
			new_num+=rmn[temp[i]]
		}
	}
	return new_num
}

func main(){
	input := "XVI"
	output := roman(input)
	fmt.Println(output)
}
