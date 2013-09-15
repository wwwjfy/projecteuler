total = 99 * 99
done = set()
for i in range(2, 11):
    this = set()
    for j in range(2, 101):
        tmp = i ** j
        if tmp > 100:
            break
        if tmp in done:
            continue
        done.add(tmp)
        for k in range(100 / j + 1, 101):
            this.add(k * j)
    total += len(this)
total -= len(done) * 99
print(total)
