#include <stdio.h>
#include <math.h>

int main(int argc, char const *argv[]) {
    int primes[150000] = {2, 3};
    int prime_count = 2;
    int i = 5;
    int sqrt_i;
    int flag;
    while (1) {
        sqrt_i = (int)sqrt((float)i);
        flag = 0;
        for (int j = 0; j < prime_count; j++) {
            if (sqrt_i < primes[j]) {
                primes[prime_count] = i;
                prime_count++;
                flag = 1;
                break;
            } else if (i % primes[j] == 0) {
                flag = 1;
                break;
            }
        }
        if (!flag) {
            primes[prime_count] = i;
            prime_count++;
        }
        i += 2;
        if (i > 2000000) {
            break;
        }
    }
    long final;
    for (int j = 0; j < prime_count; j++) {
        final += primes[j];
    }
    printf("%ld\n", final);
    return 0;
}