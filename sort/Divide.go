package sort
// 最小值最大化问题，分治贪心算法

import (
	"fmt"
)

/**
 * 把一个数组分成连续的m段，则m段中和最大的那段，最小的和是多少
 */
func Minest(array []int, m int) int {
	if m <= 0 || array == nil {
		return 0
	}
	var l, h int = 10000, 0
	for i := 0; i < len(array); i++ {
		h += array[i]
		if array[i] < l {
			l = array[i]
		}
	}
	tmp := value(array, l, h, m)
	fmt.Printf("%d\n", tmp)
	return 0
}

func value(array []int, l, r, m int) int {
	if l > r {
		return r + 1
	} else {
		mid := l + (r -l) >> 1
		if 1 == judge(array, mid, m) {
			return value(array, l, mid - 1, m)
		} else {
			return value(array, mid + 1, r, m)
		}
	}
}

func judge(array []int, mid, m int) int {
	sum := 0
	seg := 1
	for i := 0; i < len(array); i++ {
		sum += array[i]
		if sum > mid {
			sum = array[i]
			seg++
		}
	}
	if seg > m {
		return 0
	}
	return 1
}
