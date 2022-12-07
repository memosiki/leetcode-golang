from random import randint, choice


def format_list(container):
    print('[', end='')
    print(*container, sep=',', end='')
    print(']')


# a = [randint(1, 10) for _ in range(100_000)]
# a = [randint(1, 3*10**4) for i in range(3*10**4)]
a = [3*10**4 for i in range(3*10**4)]

format_list(a)
