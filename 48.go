package main

import "fmt"

func main() {
    var sum uint64
    var tmp uint64
    for i := 1; i <= 1000; i++ {
        tmp = uint64(i)
        for j := 2; j <= i; j++ {
            tmp *= uint64(i)
            tmp %= 10000000000
        }
        sum += tmp
    }
    sum %= 10000000000
    fmt.Printf("%d\n", sum)
}
