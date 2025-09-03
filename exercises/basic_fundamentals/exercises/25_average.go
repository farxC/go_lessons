package exercises

func Average(nums []float32) float32 {

	l := len(nums)

	if l == 1 {
		return nums[0]
	}

	var acc float32

	for _, v := range nums {
		acc += v
	}

	return acc / float32(l)
}
