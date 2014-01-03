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
    char numstr[7];
    size_t high, low;
    sprintf(numstr, "%u", num);
    low = 0;
    high = strlen(numstr) - 1;
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
    for (i = 1; i < 1000000; i++) {
        if (is_10_palindromic(i) && is_2_palindromic(i)) {
            sum += i;
        }
    }
    printf("%lu", sum);
    return 0;
}
