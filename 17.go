package main


// length of "one", "two", ..., "nine"
var singles = []int{3, 3, 5, 4, 4, 3, 5, 5, 4}
// length of "ten", "eleven", "twelve", ..., "nineteen"
var teens = []int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
// length of "twenty", "thirty", ..., "ninety"
var tens = []int{0, 6, 6, 5, 5, 5, 7, 6, 6}
// length of "hundred", "and"
var hundred = 7
var and = 3

func main() {
    count := 0
    one_digit := 0
    two_digit := 0

    // [1, 9]
    for i := 0; i < 9; i++ {
        count += singles[i]
    }
    one_digit = count

    // [10, 19]
    for i := 0; i <= 9; i++ {
        count += teens[i]
    }

    // [20, 99]
    for i := 1; i < 9; i++ {
        // x0
        count += tens[i]
        // others
        count += tens[i] * 9 + one_digit
    }
    two_digit = count

    // [100, 999]
    for i := 0; i < 9; i++ {
        // x00
        count += singles[i] + hundred
        // others
        count += (singles[i] + hundred + and) * 99 + two_digit
    }

    // 1000
    count += 3 + 8

    println(count)
}