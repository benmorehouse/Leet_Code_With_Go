package main

import(
	"fmt"
)

func main(){
	var array1 = []int{1,2,3,4,5}
	var array2 = []int{2,3,6,7}
	fmt.Println(merge(array1,len(array1),array2,len(array2)))
}

func merge(nums1 []int, m int, nums2 []int, n int)([]int){
    x := 0 //tracks nums1 size m
    y := 0 // tracks nums2 size n
    var newArray []int
    for i:=0;i<n+m;i++{
        if x == m{
		for i:=y;i<n;i++{
			newArray = append(newArray,nums2[i])
		}
		break
	}else if y == n{
		for i:=x;i<m;i++{
			newArray = append(newArray,nums1[i])
		}
		break
	}else{
		fmt.Println(nums1[x], " " , nums2[y])
		if nums1[x] < nums2[y]{
			newArray = append(newArray,nums1[x])
			x++
		}else if nums1[x] == nums2[y]{
			newArray = append(newArray,nums1[x])
			newArray = append(newArray,nums2[y])
			x++
			y++
		}else{
			newArray = append(newArray, nums2[y])
			y++
		}
	}
    }
	return newArray
}

/*
type LRUCache struct {
	dict map[int]int
	dictCheck map[int]bool
	size int
	maxSize int
}


func Constructor(capacity int) LRUCache {
	var this LRUCache
	if capacity <=0{
		return this
	}
	this.size = 0
	this.maxSize = capacity
	this.dict = make(map[int]int)
	this.dictCheck = make(map[int]bool)
	return this
}

func (this *LRUCache) Get(key int) int {
	if !this.dictCheck[key]{
		return -1 // means that this is not in there 
	}else{
		return this.dict[key]
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if this.size == this.maxSize{
		return
	}
	this.size++
	this.dict[key] = value
	this.dictCheck[key] = true
}


 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
