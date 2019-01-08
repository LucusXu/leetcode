package string

import (
	"fmt"
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
 * ip字符串分割
 */
func splitIp(ip string) (int, int, int, int) {
	ipArr := strings.Split(ip, ".")
	if len(ipArr) != 4 {
		return 0, 0, 0, 0
	}

	a, err:= strconv.Atoi(ipArr[0])
	if err != nil {
		return 0, 0, 0, 0
	}
	b, err:= strconv.Atoi(ipArr[1])
	if err != nil {
		return 0, 0, 0, 0
	}
	c, err:= strconv.Atoi(ipArr[2])
	if err != nil {
		return 0, 0, 0, 0
	}
	d, err:= strconv.Atoi(ipArr[3])
	if err != nil {
		return 0, 0, 0, 0
	}
	return a, b, c, d
}

func Long2Ip(long int64) string {
	if long == 0 {
		return ""
	}
	a := (long) / 256 / 256 / 256
	fmt.Printf("%d\n", a)

	b := (long - a * 256 * 256 * 256) / 256 / 256
	fmt.Printf("%d\n", b)

	c := (long - a * 256 * 256 * 256 - b * 256 * 256) / 256
	fmt.Printf("%d\n", c)

	d := (long - a * 256 * 256 * 256 - b * 256 * 256 - c * 256)
	fmt.Printf("%d\n", d)


	ip := strconv.FormatInt(a, 10) + "." + strconv.FormatInt(b, 10) + "." + strconv.FormatInt(c, 10) + "." + strconv.FormatInt(d, 10)

	x, y := Ip2Long(ip)
	fmt.Printf("%d", x, y)
	return ip
}

func Long2Ip2(long int64) string {
	if long == 0 {
		return ""
	}

	d := long % 256
	fmt.Printf("%d\n", d)

	a := (long - d) / 256 / 256 / 256
	fmt.Printf("%d\n", a)

	b := (long - 256 * 256 * 256 * a - d) / 256 / 256
	fmt.Printf("%d\n", b)

	c := (long - a * 256 * 256 * 256 - b * 256 * 256 - d)
	fmt.Printf("%d\n", c)


	ip := strconv.FormatInt(a, 10) + "." + strconv.FormatInt(b, 10) + "." + strconv.FormatInt(c, 10) + "." + strconv.FormatInt(d, 10)

	x, y := Ip2Long(ip)
	fmt.Printf("%d %d \n", x, y)
	return ip
}
