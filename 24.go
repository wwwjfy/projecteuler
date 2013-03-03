package main

import "fmt"

var offset int = 999999

func main() {
    var factorial [10]int
    var quotient int
    var arrange [10]int
    var length int = 0
    var nums [10]int
    for i := 0; i < 10; i++ {
        nums[i] = i
    }
    factorial[1] = 1
    for i := 2; i < 10; i++ {
        factorial[i] = i * factorial[i-1]
    }
    for i := 9; i > 0; i-- {
        quotient = offset / factorial[i]
        arrange[length] = quotient
        length++
        offset -= quotient * factorial[i]
        if offset == 0 {
            break
        }
    }
    for i := 0; i < 10; i++ {
        fmt.Printf("%d", nums[arrange[i]])
        for j := arrange[i]; j < 9 - i; j++ {
            nums[j] = nums[j+1]
        }
    }
}