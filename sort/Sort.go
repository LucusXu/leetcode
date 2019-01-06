package sort

import (
	// "fmt"
)

/**
 * 冒泡排序
 */
func Bubble(array []int, length int)  {
	if length <= 0 || array == nil {
		return
	}
	for i := length -1; i > 0; i-- {
		for j := 0; j < i; j++ {
			// fmt.Printf("%d, %d\n", i, j)
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

/**
 * 快速排序
 */
func Qsort(array []int, length int)  {
	if length <= 0 || array == nil {
		return
	}
	// fmt.Printf("%v", array, length)
	var i int = 0
	var j int = length - 1
	var p int = length / 2
	// fmt.Printf("%v", i, j, p)

	for i != j {
		if array[i] < array[p] {
			i++
		} else {
			array[i], array[p] = array[p], array[i]
			p = i
			// fmt.Printf("%d, %d\n", i, p)
		}
		if array[p] < array[j] {
			j--
		} else {
			array[j], array[p] = array[p], array[j]
			p = j
		}
		// fmt.Printf("%v", i, j ,p)
		// fmt.Printf("%v", array)
	}
	Qsort(array[:p], p)
	Qsort(array[p + 1:], length - p - 1)
}
