package base

import "fmt"

func Mono(nums []int) bool {
	size:=len(nums)

	if size<=1 {
		return false
	}

	inc := true
	dec := true
 
	for i:=0; i<size-1; i++ {

		if nums[i]>=nums[i+1] {
		return inc = false
		}
		if nums[i]<=nums[i+1] {
			return dec = false
		}

	}
	return inc || dec
}