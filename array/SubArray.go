package array
// 和最大的最短子数组问题

func ShortestSubArray(array []int, n int) []int {
	if array == nil || n <= 0 {
		return nil
	}
	left, sum := 0, 0
	size := len(array)
	var ret []int
	for i := 0;i < size;i++ {
		sum = sum + array[i]
		for ;left <= i && sum >= n; {
			sum = sum - array[left]
			left++
			ret = array[left:]
		}
	}
	return ret
}