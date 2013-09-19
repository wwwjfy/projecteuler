package main

import "fmt"

var results []int = make([]int, 0, 10)

func check(number [5]int, result int, display int) {
    mark := [10]bool{}
    mark[0] = true
    for i := 0; i < 5; i++ {
        mark[number[i]] = true
    }
    mid := result
    var tmp int
    for mid > 0 {
        tmp = mid % 10
        if mark[tmp] {
            return
        }
        mark[tmp] = true
        mid /= 10
    }
    switch (display) {
    case 1:
        fmt.Printf("%d * %d%d%d%d = %d\n", number[0], number[1], number[2], number[3], number[4], result)
    case 2:
        fmt.Printf("%d%d * %d%d%d = %d\n", number[0], number[1], number[2], number[3], number[4], result)
    default:
        break
    }
    for i := 0; i < len(results); i++ {
        if result == results[i] {
            return
        }
    }
    results = append(results, result)
}

func validate(number[5] int) {
    if number[0] * number[1] <= 9 {
        check(number, number[0] * (number[1] * 1000 + number[2] * 100 + number[3] * 10 + number[4]), 1)
    }
    if number[0] * number[2] <= 9 {
        check(number, (number[0] * 10 + number[1]) * (number[2] * 100 + number[3] * 10 + number[4]), 2)
    }
}

func fill(number [5]int, pos int) {
    if (pos == 5) {
        validate(number)
        return
    }
    mark := [10]bool{}
    for i := 0; i < pos; i++ {
        mark[number[i]] = true
    }
    for i := 1; i <= 9; i++ {
        if !mark[i] {
            number[pos] = i
            fill(number, pos+1)
        }
    }
}

func main() {
    var number [5]int
    fill(number, 0)
    total := 0
    for i := 0; i < len(results); i++ {
        total += results[i]
    }
    fmt.Printf("%d\n", total)
}
