package main

import "fmt"

var max int = 0

func calc(numbers [400]int, index func(int, int) (int, int, bool)) {
    var current int
    var count int
    for i := 0; i < 20; i++ {
        // filter 0
        current = 1
        count = 0
        for j := 0; j < 20; j++ {
            now, first, err := index(i, j)
            if err {
                break
            }
            if numbers[now] == 0 {
                current = 1
                count = 0
                continue
            }
            current *= numbers[now]
            if count < 4 {
                count ++
            } else {
                current /= numbers[first]
            }
            if current > max && count == 4 {
                max = current
            }
        }
    }
}

func by_row(i, j int) (int, int, bool) {
    return i * 20 + j, i * 20 + j - 4, false
}

func by_col(i, j int) (int, int, bool) {
    return j * 20 + i, (j - 4) * 20 + i, false
}

func by_diagonal_left_top(i, j int) (int, int, bool) {
    if i < 4 || i < j {
        return 0, 0, true
    }
    return (i - j) * 20 + j, (i - (j - 4)) * 20 + (j - 4), false
}

func by_diagonal_right_bottom(i, j int) (int, int, bool) {
    if i == 0 || i > 16 || 19 - i < j {
        return 0, 0, true
    }
    return (19 - j) * 20 + i + j, (19 - (j - 4)) * 20 + i + j - 4, false
}

func by_diagonal_right_top(i, j int) (int, int, bool) {
    if i < 4 || i < j {
        return 0, 0, true
    }
    return j * 20 + (19 - i) + j, (j - 4) * 20 + (19 - i) + j - 4, false
}

func by_diagonal_left_bottom(i, j int) (int, int, bool) {
    if i == 0 || i > 16 || 19 - i < j {
        return 0, 0, true
    }
    return (i + j) * 20 + j, (i + j - 4) * 20 + j - 4, false
}

func main() {
    var numbers [400]int
    // read numbers
    for i := 0; i < 400; i++ {
        fmt.Scanf("%d", &numbers[i])
    }
    // calculate by lines
    calc(numbers, by_row)
    // calculate by columns
    calc(numbers, by_col)
    // calculate by diagonal lines
    calc(numbers, by_diagonal_left_top)
    calc(numbers, by_diagonal_right_bottom)
    calc(numbers, by_diagonal_right_top)
    calc(numbers, by_diagonal_left_bottom)
    fmt.Printf("%d\n", max)
}