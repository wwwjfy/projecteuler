package main

var days []int = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func is_leap_year(year int) bool {
    if (year % 100 == 0 && year % 400 == 0) || 
       (year % 100 != 0 && year % 4 == 0) {
        return true
    }
    return false
}

func days_of_month(year, month int) int {
    if month != 2 {
        return days[month]
    }
    if is_leap_year(year) {
        return 29
    }
    return 28
}

func main() {
    count := 0
    weekday := 0
    if is_leap_year(1900) {
        weekday = (weekday + 366) % 7
    } else {
        weekday = (weekday + 365) % 7
    }
    for i := 1901; i <= 2000; i++ {
        for j := 1; j <= 12; j++ {
            if weekday == 6 {
                count++
            }
            weekday = (weekday + days_of_month(i, j)) % 7
        }
    }
    println(count)
}
