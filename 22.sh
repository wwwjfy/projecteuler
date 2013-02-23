#!/bin/bash

sed 's/,/\
/g
s/"//g' 22.txt | sort | awk 'BEGIN {
    for (i = 65; i <= 90; i++) {
        ord[sprintf("%c", i)] = i - 64
    }
    product_sum = 0
}
{
    product = 0
    for (i = 1; i <= length($0); i++) {
        product += ord[substr($0, i, 1)]
    }
    product *= NR
    product_sum += product
}
END {
    print product_sum
}
'