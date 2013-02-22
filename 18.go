package main

import "fmt"

const incre int = 10

func more_row(rows *[][]int, length, num int) int {
    if len(*rows) == length {
        var old_rows [][]int
        old_rows = *rows
        *rows = make([][]int, length + incre)
        copy(*rows, old_rows)
    }
    (*rows)[length] = make([]int, num)
    return length + 1
}

func bigger(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    var rows [][]int
    length := 0
input:
    for i := 0; ; i++ {
        length = more_row(&rows, length, i + 1)
        for j := 0; j < i + 1; j++ {
            _, err := fmt.Scanf("%d", &rows[i][j])
            if err != nil {
                break input
            }
        }
    }
    length--

    // skip the bottom line
    for i := length - 2; i >= 0; i-- {
        for j := 0; j < i + 1; j++ {
            rows[i][j] += bigger(rows[i+1][j], rows[i+1][j+1])
        }
    }
    println(rows[0][0])
}