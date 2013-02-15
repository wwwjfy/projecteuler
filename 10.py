primes = [2, 3]

i = 5
while True:
    sqrt_i = int(i ** 0.5)
    for prime in primes:
        if sqrt_i < prime:
            primes.append(i)
            break
        elif i % prime == 0:
            break
    else:
        primes.append(i)

    i += 2
    if i >= 2000000:
        break

print sum(primes)
