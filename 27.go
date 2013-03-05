package main

import "math"

func is_prime(num int) bool {
    if num <= 0 {
        return false
    }
    if num == 2 {
        return true
    }
    var sqrt_num int = int(math.Sqrt(float64(num)))
    for i := 2; i <= sqrt_num; i++ {
        if num % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var a, b int
    var count, max int
    for i := -999; i < 1000; i++ {
        for j := 0; j < 1000; j++ {
            count = 0
            for ; ; count++ {
                if !is_prime(count * count + i * count + j)  {
                    count--
                    break
                }
            }
            if count > max {
                max = count
                a = i
                b = j
            }
        }
    }
    // println(a, b)
    // println(max)
    println(a * b)
}