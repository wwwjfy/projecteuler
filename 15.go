package main

import "fmt"

func main() {
    const num = 22
    var col [num]int64
    var row [num]int64
    for i := 0; i < num; i++ {
        col[i] = 1
        row[i] = 1
    }
    for i := 2; i < num; i++ {
        for j := 1; j < i - 1; j++ {
            row[j] = row[j - 1] + row[j]
            col[j] = col[j - 1] + col[j]
        }
        row[i - 1] = row[i - 2] + col[i - 2]
        col[i - 1] = row[i - 1]
        fmt.Printf("%d: %d\n", i - 1, row[i - 1])
    }
}