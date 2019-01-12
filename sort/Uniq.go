package sort

import (
	"fmt"
	"math"
)

/**
 * 二分查找算法循环实现
 */
func UniqCount(array []int, length int) int {
	if length <= 0 || array == nil {
		return 0
	}
	var l int = 0
	var h int = length - 1
	fmt.Printf("%d,%d\n", l, h)
	var count int = 0
	for ; l <= h; {
		valuel := math.Abs(float64(array[l]))
		valueh := math.Abs(float64(array[h]))
		if valuel == valueh {
			// fmt.Printf("%v\n", l, h)
			count++
			for ; l <= h; {
				v := math.Abs(float64(array[l]))
				if valuel == v {
					l++
				} else {
					break
				}
			}
			for ; l <= h; {
				v := math.Abs(float64(array[h]))
				if valueh == v {
					h--
				} else {
					break
				}
			}
		} else if valuel > valueh {
			count++
			for ; l<=h ; {
				v := math.Abs(float64(array[l]))
				if valuel == v {
					l++
				} else {
					break
				}
			}
		} else {
			count++
			for ; l<=h; {
				v := math.Abs(float64(array[h]))
				if valueh == v {
					h--
				} else {
					break
				}
			}
		}
	}

	return count
}
