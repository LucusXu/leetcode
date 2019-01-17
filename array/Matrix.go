package array
// 二维数组查找

func Search(array [][]int, m, n, t int) bool {
	if array == nil || m <= n || n <= 0 {
		return false
	}

	i, j := 0, n - 1
	for ; i < m && j >= 0; {
		if array[i][j] == t {
			return true
		} else if array[i][j] > t {
			j--
		} else {
			i++
		}
	}
	return false
}