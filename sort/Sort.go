package sort

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
	for i := length -1; i > 0; i-- {
		for j := 0; j < i; j++ {
			// fmt.Printf("%d, %d\n", i, j)
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
