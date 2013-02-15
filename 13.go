package main

import "fmt"

func main() {
    var strs [100]string
    var temp int
    var length int
    var result [5000]byte
    for i := 0; i < 100; i++ {
        fmt.Scanf("%s", &strs[i])
    }
    temp = 0
    for i := 49; i >= 0; i-- {
        for j := 0; j < 100; j++ {
            temp += int(strs[j][i]) - 48
        }
        result[length] = byte(temp % 10 + 48)
        length++
        temp /= 10
    }
    for temp > 0 {
        length++
        result[length-1] = byte(temp % 10 + 48)
        temp /= 10
    }
    for i := length - 1; i >= length - 1 - 9; i-- {
        fmt.Printf("%c", result[i])
    }
}