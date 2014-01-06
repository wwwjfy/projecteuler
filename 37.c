#include <stdio.h>
#include <math.h>
#include <stdbool.h>
#include <string.h>

bool is_prime(int num) {
    if (num == 1) {
        return false;
    }
    int s = (int)sqrt(num);
    int i;
    for (i = 2; i <= s; i++) {
        if (num % i == 0) {
            return false;
        }
    }
    return true;
}

bool is_left_truncatble(int num) {
    char str[100];
    int digits;
    sprintf(str, "%d", num);
    digits = (int)pow(10, (double)strlen(str) - 1);
    while (digits > 1) {
        num %= digits;
        if (!is_prime(num)) {
            return false;
        }
        digits /= 10;
    }
    return true;
}

bool is_right_truncatble(int num) {
    while ((num /= 10)) {
        if (!is_prime(num)) {
            return false;
        }
    }
    return true;
}

int main() {
    int i = 11;
    int count = 0;
    int sum = 0;
    while (count < 11) {
        if (is_prime(i) && is_left_truncatble(i) && is_right_truncatble(i)) {
            printf("%d\n", i);
            sum += i;
            count++;
        }
        i += 2;
    }
    printf("The result: %d\n", sum);
    return 0;
}
