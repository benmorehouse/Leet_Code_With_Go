package main

import(
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type listNode struct{
	data int
	next *listNode
}

func addTwoNumbers(l1 *listNode, l2 *listNode)(int){
	num1 := getNumber(l1)
	num2 := getNumber(l2)
	return num1+num2
}

func getNumber(list *listNode)(int){
	tenths := 1
	temp := 0

	for list != nil{
		temp += list.data * tenths
		tenths*=10
		list = list.next
	}
	return temp
}

func main(){

	var mylist = listNode{data: 1}
	var mylist2 = listNode{data: 2}
	var mylist3 = listNode{data: 3}
	*mylist.next = mylist2
	*mylist2.next = mylist3

	var mylist4 = listNode{data: 3}
	var mylist5 = listNode{data: 3}
	var mylist6 = listNode{data: 1}
	*mylist4.next = mylist5
	*mylist5.next = mylist6

	fmt.Println(addTwoNumbers(&mylist,&mylist4))
}




