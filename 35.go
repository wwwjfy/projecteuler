package main

import "fmt"
import "math"

var total = 5 // 2, 3, 5, 7, 11
var multiplier = []int{1, 10, 100, 1000, 10000, 100000}

func checkIfPrime(primes []int, num int) bool {
    low, high := 0, len(primes) - 1
    var mid int
    for low <= high {
        mid = low + (high - low) / 2
        if primes[mid] > num {
            high = mid - 1
        } else if primes[mid] < num {
            low = mid + 1
        } else {
            return true
        }
    }
    return false
}

func checkCircular(primes []int, num int) {
    numbers := [6]int{}
    i := 0
    max := 0
    for num > 0 {
        numbers[i] = num % 10
        if numbers[i] % 2 == 0 {
            return
        }
        num /= 10
        if max < numbers[i] {
            max = numbers[i]
        }
        i++
    }
    if numbers[i - 1] != max {
        return
    }
    var other int
    for j := 1; j < i; j++ {
        other = 0
        for k := j; k < j + i; k++ {
            other += numbers[k % i] * multiplier[k - j]
        }
        if !checkIfPrime(primes, other) {
            return
        }
    }
    fmt.Printf("%d %d\n", other, i)
    total += i
}

func main() {
    primes := []int{2, 3, 5, 7, 11, 13}
    var flag bool
    var sqrt int
    for i := 15; i < 1000000; i += 2 {
        flag = false
        sqrt = int(math.Sqrt(float64(i)))
        for j := 1; j < len(primes); j++ {
            if primes[j] > sqrt {
                break
            }
            if i % primes[j] == 0 {
                flag = true
                break
            }
        }
        if !flag {
            primes = append(primes, i)
            checkCircular(primes, i)
        }
    }
    fmt.Printf("Total primes: %d\n", len(primes))
    fmt.Printf("result: %d\n", total)
}
