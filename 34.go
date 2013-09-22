package main

import "fmt"

func main() {
    factorial := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
    var tmp int
    var sum int
    for i := 3; i < factorial[9] * 6 + factorial[2]; i++ {
        tmp = i
        sum = 0
        for tmp > 0 {
            sum += factorial[tmp % 10]
            tmp /= 10
        }
        if sum == i {
            fmt.Printf("%d\n", i)
        }
    }
}
