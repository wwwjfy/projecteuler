package main

import "fmt"

func main() {
    nums := [10000]bool{}
    sum := 0
    var num, divisors_sum, last int
    for i := 1; i < 10000; i++ {
        if nums[i] {
            continue
        }
        num = i
        last = i
        for {
            if nums[num] {
                break
            }
            divisors_sum = 1
            for j := 2; j < num; j++ {
                if num % j == 0 {
                    divisors_sum += j
                }
            }
            nums[num] = true
            if divisors_sum == num || divisors_sum >= 10000 {
                break
            }
            if divisors_sum == last {
                fmt.Printf("%d %d\n", num, last)
                sum += num
                sum += last
                break
            }
            last = num
            num = divisors_sum
        }
    }
    println(sum)
}