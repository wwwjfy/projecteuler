package main

import "fmt"

func main() {
    var maxcount int64 = 0
    var count int64
    var num int64
    var maxnum int64
    var i int64
    for i = 500000; i < 1000000; i++ {
        count = 0
        num = i
        for {
            if num & 1 == 1 {
                num = 3 * num + 1
            } else {
                num /= 2
            }
            count++
            if num == 1 {
                break
            }
        }
        if count > maxcount {
            maxnum = i
            maxcount = count
        }
    }
    fmt.Printf("%d: %d", maxnum, maxcount)
}
