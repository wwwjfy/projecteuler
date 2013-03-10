package main

func main() {
    sum := 1
    tmp := 1
    for i := 1; i <= 500; i++ {
        sum += 4 * tmp + i * 2 * 10
        tmp += i * 2 * 4
    }
    println(sum)
}