package main

import "fmt"
import "math"

func main() {
	primes := [150000]int{2, 3}
	prime_count := 2
	now := 5
	var sqrt_now int
	for {
		sqrt_now = int(math.Sqrt(float64(now)))
		flag := false
		for i := 0; i < prime_count; i++ {
			if sqrt_now < primes[i] {
				primes[prime_count] = now
				prime_count++
				flag = true
				break
			} else if now%primes[i] == 0 {
				flag = true
				break
			}
		}
		if !flag {
			primes[prime_count] = now
			prime_count++
		}
		now += 2

		if now > 2000000 {
			break
		}
	}

	var final int64 = 0
	for i := 0; i < prime_count; i++ {
		final += int64(primes[i])
	}
	fmt.Printf("%d\n", final)
}
