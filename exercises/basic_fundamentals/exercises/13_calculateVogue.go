package exercises

import (
	"math"
	"strconv"
)

func CalculateVogue(nums []int) (int, string) {
	count := make(map[string]int)
	for i := range nums {
		strRep := strconv.Itoa(nums[i])
		_, ok := count[strRep]
		if ok {
			count[strRep] += 1
		} else {
			count[strRep] = 1
		}
	}
	maxVal := math.MinInt
	var vogueKey string
	for k, v := range count {
		if v > maxVal {
			maxVal = v
			vogueKey = k
		}
	}
	return maxVal, vogueKey
}
