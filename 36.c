#include <stdio.h>
#include <string.h>
#include <stdbool.h>

bool is_2_palindromic(unsigned num) {
    unsigned high = 1 << (sizeof(unsigned) * 8 - 1);
    unsigned low = 1;
    while (!(num & high) && (high >>= 1));
    while (high > low) {
        if (((num & high) && !(num & low)) ||
            (!(num & high) && (num & low))) {
            return false;
        }
        high >>= 1;
        low <<= 1;
    }
    return true;
}

bool is_10_palindromic(unsigned num) {
    char numstr[6];
    char tmp;
    size_t high = 0, low = 0;
    tmp = num % 10;
    num /= 10;
    while (num || tmp) {
        numstr[high] = tmp;
        high++;
        tmp = num % 10;
        num /= 10;
    }
    high--;
    while (high > low) {
        if (numstr[high] != numstr[low]) {
            return false;
        }
        high--;
        low++;
    }
    return true;
}

int main(void) {
    long unsigned sum = 0;
    unsigned i;
    // 0 at the end of binary for even numbers
    for (i = 1; i < 1000000; i += 2) {
        if (is_10_palindromic(i) && is_2_palindromic(i)) {
            sum += i;
        }
    }
    printf("%lu", sum);
    return 0;
}
