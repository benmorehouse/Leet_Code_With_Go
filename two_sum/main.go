/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/


package main

import(
	"fmt"
)
// o(n) runtime worst case scenario for both
func twoSums(nums []int, target int)([]int){
	var solution []int
	if target < 0{
		return solution
	}else if len(nums) == 0{
		return solution
	}
	// need to go through and add the subsets of all the numbers
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums);j++{
			if i==j{
				continue
			}else if nums[i] + nums[j] == target{
				solution = append(solution,nums[i])
				solution = append(solution,nums[j])
				return solution
			}
		}
	}
	return solution
}


func main(){
	nums := []int{8,3,4,7,9,1}
	target := 9
	answer := twoSums(nums, target)
	fmt.Println("answer is:",answer)
}
