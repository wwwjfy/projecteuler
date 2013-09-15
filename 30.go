package main

func main() {
    MAX := 9 * 9 * 9 * 9 * 9 * 5
    pre_results := [11]int{}
    var quotient, remainder int
    var result int
    var total int
    for i := 1; i <= 9; i++ {
        pre_results[i] = i * i * i * i * i
    }
    for i := 2; i <= MAX; i++ {
        quotient, remainder = i / 10, i % 10
        result = 0
        for {
            result += pre_results[remainder]
            if quotient == 0 {
                break
            }
            quotient, remainder = quotient / 10, quotient % 10
        }
        if result == i {
            total += result
        }
    }
    print(total)
}
