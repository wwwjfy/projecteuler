def one_child(length):
    total = 0
    for i in range(1, length):
        to_test = split(i)
        if to_test:
            total += verify(to_test, len(str(i)))
    print total


def split(num):
    splited = set()
    num = str(num)
    current = 1
    length = len(num)
    while current <= length:
        for i in range(0, length - current + 1):
            now = num[i:i+current]
            splited.add(int(now))
        current += 1
    return list(splited)


def verify(to_test, length):
    flag = False
    for i in to_test:
        if i % length == 0:
            if flag:
                return 0
            else:
                flag = True
    if flag:
        return 1
    else:
        return 0

# TODO: optimize split algorithm
one_child(1000)
# print split(5671)