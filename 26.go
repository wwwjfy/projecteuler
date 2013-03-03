package main

func main() {
    longest := 0
    d := 1
    for i := 2; i < 1000; i++ {
        m := make([]int, 1000)
        divider := 10
        length := 0
        for {
            remainder := divider % i
            if remainder == 0 {
                break
            }
            if m[remainder] == 0 {
                length++
                m[remainder] = length
            } else {
                tmp := length - m[remainder] + 1
                if tmp > longest {
                    longest = tmp
                    d = i
                }
                break
            }
            divider = remainder * 10
        }
    }
    println(d)
}