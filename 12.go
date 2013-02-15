package main

// import "fmt"
import "math"

func main() {
    triangular := 28
    incre := 8
    num := 0
    sqrt := 0
    for {
        sqrt = int(math.Sqrt(float64(triangular)))
        for i := 2; i < sqrt; i++ {
            if triangular % i == 0 {
                num++
            }
        }
        if num > 250 {
            break
        }
        num = 0
        triangular += incre
        incre++
    }
    println(triangular)
}