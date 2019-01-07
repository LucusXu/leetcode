package string

import (
	// "fmt"
	"strings"
	"strconv"
)

func Ip2Long(ip string) (int64, int64) {
	a, b, c, d := splitIp(ip)
	if a == 0 {
		return 0, 0
	}
	x := (a << 24) | (b << 16) | (c << 8) | d
	y := a * 256 * 256 * 256 + b * 256 * 256 + c * 256 + d
	return int64(x), int64(y)
}

/**
 * 二分查找算法循环实现
 */
func splitIp(ip string) (int, int, int, int) {
	ipArr := strings.Split(ip, ".")
	if len(ipArr) != 4 {
		return 0, 0, 0, 0
	}

	a, err:= strconv.Atoi(ipArr[3])
	if err != nil {
		return 0, 0, 0, 0
	}
	b, err:= strconv.Atoi(ipArr[2])
	if err != nil {
		return 0, 0, 0, 0
	}
	c, err:= strconv.Atoi(ipArr[1])
	if err != nil {
		return 0, 0, 0, 0
	}
	d, err:= strconv.Atoi(ipArr[0])
	if err != nil {
		return 0, 0, 0, 0
	}
	return a, b, c, d
}
