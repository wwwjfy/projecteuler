/*

The idea is for all bases who are results of others,
use the very original base to count the exponents.

For example, 4^51 will be counted as 2^102, which is
also the result of 8^34, won't be counted multiple times.

In the meanwhile, 4, 8, ... need not to be in, again.

*/
package main

import "fmt"

func isIn(dup []int, elem int) bool {
    for i := 0; i < len(dup); i++ {
        if dup[i] == elem {
            return true
        }
    }
    return false
}

func main() {
    var tmp, count int
    var dup = []int{}
    var total = 99 * 99
    var elems = []int{}
    var elem_count int
    for i := 2; i <= 10; i++ {
        if isIn(dup, i) {
            continue
        }
        tmp = i * i
        count = 2
        elems = make([]int, 100)
        elem_count = 0
        for {
            if tmp > 100 {
                break
            }
            dup = append(dup, tmp)
            total -= 99
            for j := 100 / count + 1; j <= 100; j++ {
                if isIn(elems, j * count) {
                    continue
                }
                elems = append(elems, j * count)
                elem_count++
            }
            tmp *= i
            count++
        }
        total += elem_count
    }
    fmt.Printf("%d\n", total)
}
