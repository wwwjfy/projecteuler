package main

import "fmt"

func main() {
    var CAND = []int{1, 2, 5, 10, 20, 50, 100, 200}
    var count = [201]int{}
    count[0] = 1

    for i := 0; i < len(CAND); i++ {
        for j := CAND[i]; j <= 200; j++ {
            count[j] += count[j - CAND[i]]
        }
    }
    fmt.Printf("%d", count[200])
}
