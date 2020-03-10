package main
import "fmt"

func sortByBits(arr []int) []int {
    var bitmap = make(map[int][]int, 0)
    // map缓存
    for i := 0; i < len(arr); i++ {
        bits := calBits(arr[i])
        if _, ok := bitmap[bits]; !ok {
            tmp := make([]int, 0)
            tmp = append(tmp, arr[i])
            bitmap[bits] = tmp
        } else {
            bitmap[bits] = append(bitmap[bits], arr[i])
            /*
            for j := len(bitmap[bits]) - 1;j > 0; j-- {
                if bitmap[bits][j] < bitmap[bits][j - 1] {
                    bitmap[bits][j], bitmap[bits][j - 1] = bitmap[bits][j - 1], bitmap[bits][j]
                }
            }*/
        }
    }
    fmt.Printf("%v", bitmap)
    // 输出
    var res []int
    for key, value := range bitmap {
        fmt.Printf("%v%v", key, value)
        res = append(res, value...)
    }
    fmt.Printf("%v", res)
    return res
}

func calBits(n int) int {
    sum := 0
    for n > 0 {
        if n & 1 == 1 {
            sum++
        }
        n = n >> 1
    }
    return sum
}

func main() {
    var A = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
    sortByBits(A)
}
