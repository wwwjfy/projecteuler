package main

func main() {
    var nums [28124]bool
    var abundant [28124]int
    var abundant_count int = 0
    var divisors_sum int
    for i := 12; i < 28124; i++ {
        divisors_sum = 1
        for j := 2; j < i; j++ {
            if i % j == 0 {
                divisors_sum += j
            }
        }
        if divisors_sum > i {
            abundant[abundant_count] = i
            abundant_count++
        }
    }
    for i := 0; i < abundant_count; i++ {
        for j := 0; j < abundant_count; j++ {
            if abundant[i] + abundant[j] < 28124 {
                nums[abundant[i] + abundant[j]] = true
            }
        }
    }
    sum := 0
    for i := 1; i < 28124; i++ {
        if !nums[i] {
            sum += i
        }
    }
    println(sum)
}
