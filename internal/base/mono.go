package base

func Mono(nums []int) bool {
	size := len(nums)

	if size <= 1 {
		return true
	}

	inc, dec := true, true

	for i := 0; i < size-1; i++ {

		if nums[i] > nums[i+1] {
			inc = false
		}
		if nums[i] < nums[i+1] {
			dec = false
		}

	}
	return inc || dec
}
