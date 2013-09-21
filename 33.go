package main

import "fmt"

var numerator, denominator int = 1, 1

func check(numerator1, denominator1, numerator2, denominator2 int) {
    if float32(numerator1) / float32(denominator1) == float32(numerator2) / float32(denominator2) && numerator2 < denominator2 {
        fmt.Printf("%d / %d = %d / %d\n", numerator1, denominator1, numerator2, denominator2)
        numerator *= numerator1
        denominator *= denominator1
    }
}

func validate(number [3]int) {
    check(number[0] * 10 + number[1], number[1] * 10 + number[2], number[0], number[2])
    check(number[0] * 10 + number[2], number[1] * 10 + number[2], number[0], number[1])
    check(number[0] * 10 + number[1], number[2] * 10 + number[0], number[1], number[2])
}

func fill(number [3]int, pos int) {
    if pos == 3 {
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

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}

func main() {
    var number [3]int
    fill(number, 0)
    fmt.Printf("The product is %d / %d\n", numerator, denominator)
    fmt.Printf("The result is %d\n", denominator / (gcd(numerator, denominator)))
}
