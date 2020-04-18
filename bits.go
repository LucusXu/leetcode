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

func findMinSum(arr []int, sum int) int {
    l := len(arr)
    if l == 0 {
        return 0
    }
    if l == 1 && arr[0] >= sum {
        return 1
    }
    // 返回值
    res := 0
    p := 0
    q := p
    // 临时和
    tmpSum := arr[0]
    for p < l && q < l {
        if tmpSum >= sum {
            if res == 0 || q-p+1 < res {
                res = q - p + 1
            }
            tmpSum = tmpSum - arr[p]
            p++
        } else {
            q++
            if q >= l {
                break
            }
            tmpSum = tmpSum + arr[q]
        }
    }
    return res
}

func main() {
    var A = []int{4, 5, 1, 2, 6}
    ret := findMinSum(A, 1)
    fmt.Println("%d", ret)
}
