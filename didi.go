package main

import "fmt"

func change(numbers string) string {
	l := len(numbers)
	if l < 2 {
		return numbers
	}
	var num = []byte(numbers)

	for i := 0; i < l - 1; i++ {
		max := -1
		for j := i + 1; j < l; j++ {
			if num[i] == num[j] {
				continue
			}
			if num[j] > num[i] {
				max = j
			}
		}
		if max != -1 {
			num[i], num[max] = num[max], num[i]
			break
		}
	}
	fmt.Println("%d", string(num))
	return string(num)
}

func greedy(numbers string) string {
	// 贪心算法,记录每个数字出现的最后位置
	var cache = make([]int, 10)
	fmt.Println("%d", len(cache))
	l := len(numbers)
	str := []byte(numbers)
	for i := 0; i < l; i++ {
		tmp := int(str[i] - '0')
		fmt.Println("%d", tmp, i)
		cache[tmp] = i
	}
	// 找到比它大的最大位置, 交换
	for i := 0; i < l; i++ {
		tmp := int(str[i] - '0')
		flag := false
		for j := 9; j > tmp; j-- {
			idx := cache[j]
			if idx > i {
				str[i], str[idx] = str[idx], str[i]
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	fmt.Println("%d\n", string(str))
	return string(str)
}

func maximumSwap(num int) int {
	// 转换成数组
	arr := int2Arr(num)
	fmt.Println("%d", arr)
	// 贪心算法记录每个数字最后出现的位置
	var cache = make([]int, 10)
	for i := 0; i < len(arr); i++ {
		cache[arr[i]] = i
	}
	for i := 0; i < len(arr); i++ {
		flag := false
		tmp := arr[i]
		for j := 9; j > tmp; j-- {
			idx := cache[j]
			if idx > i {
				arr[i], arr[idx] = arr[idx], arr[i]
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	return arr2Int(arr)
}

func arr2Int(arr []int) int {
	var res = 0
	for i := 0; i < len(arr); i++ {
		res = res * 10 + arr[i]
	}
	return res
}

func int2Arr(num int) []int {
	var res []int
	if num == 0 {
		res = append(res, 0)
	} else {
		for num > 0 {
			res = append(res, num % 10)
			num = num / 10
		}
	}
	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
	}
	return res
}


func main() {
	// var numbers = "2736"
	// change(numbers)
	// greedy(numbers)
	maximumSwap(2736)
}
