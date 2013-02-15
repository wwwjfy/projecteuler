#include <stdio.h>

int max = 0;

typedef struct {
    unsigned int current;
    unsigned int first;
    int error:1;
} IndexValue;

void calc(int *numbers, IndexValue(*index)(int, int)) {
    int current, count;
    IndexValue index_value;
    for (int i = 0; i < 20; i++) {
        current = 1;
        count = 0;
        for (int j = 0; j < 20; j++) {
            index_value = index(i, j);
            if (index_value.error) {
                break;
            }
            if (numbers[index_value.current] == 0) {
                current = 1;
                count = 0;
                continue;
            }
            current *= numbers[index_value.current];
            if (count < 4) {
                count++;
            } else {
                current /= numbers[index_value.first];
            }
            if (current > max && count == 4) {
                max = current;
            }
        }
    }
    return;
}

IndexValue by_row(int i, int j) {
    return (IndexValue){i * 20 + j, i * 20 + j - 4, 0};
}

IndexValue by_col(int i, int j) {
    return (IndexValue){j * 20 + i, (j - 4) * 20 + i, 0};
}

IndexValue by_diagonal_left_top(int i, int j) {
    if (i < 4 || i < j) {
        return (IndexValue){0, 0, 1};
    }
    return (IndexValue){(i - j) * 20 + j, (i - (j - 4)) * 20 + (j - 4), 0};
}

IndexValue by_diagonal_right_bottom(int i, int j) {
    if (i == 0 || i > 16 || 19 - i < j) {
        return (IndexValue){0, 0, 1};
    }
    return (IndexValue){(19 - j) * 20 + i + j, (19 - (j - 4)) * 20 + i + j - 4, 0};
}

IndexValue by_diagonal_right_top(int i, int j) {
    if (i < 4 || i < j) {
        return (IndexValue){0, 0, 1};
    }
    return (IndexValue){j * 20 + (19 - i) + j, (j - 4) * 20 + (19 - i) + j - 4 , 0};
}

IndexValue by_diagonal_left_bottom(int i, int j) {
    if (i == 0 || i > 16 || 19 - i < j) {
        return (IndexValue){0, 0, 1};
    }
    return (IndexValue){(i + j) * 20 + j, (i + j - 4) * 20 + j - 4, 0};
}

int main(int argc, char *argv[]) {
    int numbers[20][20];
    for (int i = 0; i < 20; i++) {
        for (int j = 0; j < 20; j++)
        {
            scanf("%d", &numbers[i][j]);
        }
    }
    calc(numbers, by_row);
    calc(numbers, by_col);
    calc(numbers, by_diagonal_left_top);
    calc(numbers, by_diagonal_right_bottom);
    calc(numbers, by_diagonal_right_top);
    calc(numbers, by_diagonal_left_bottom);
    printf("%d\n", max);
    return 0;
}