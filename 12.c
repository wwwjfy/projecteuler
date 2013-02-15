#include <stdio.h>
#include <math.h>

int main(int argc, char const *argv[]) {
    long triangular = 28;
    unsigned incre = 8;
    unsigned num = 8;
    unsigned sqrted = 0;
    while (1) {
        sqrted = (int)sqrt((double)triangular);
        for (int i = 2; i < sqrted; i++) {
            if (triangular % i == 0) {
                num++;
            }
        }
        if (num > 250) {
            break;
        }
        num = 0;
        triangular += incre;
        incre++;
    }
    printf("%ld\n", triangular);
    return 0;
}