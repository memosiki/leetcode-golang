def product(*args, repeat=1):
    # product('ABCD', 'xy') --> Ax Ay Bx By Cx Cy Dx Dy
    # product(range(2), repeat=3) --> 000 001 010 011 100 101 110 111
    pools = [tuple(pool) for pool in args] * repeat
    result = [[]]
    for pool in pools:
        result = [x + [y] for x in result for y in pool]
    for prod in result:
        # yield prod[0]+prod[1]
        yield tuple(prod)


def letterCombinations(digits: str):
    dial, ans = {'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}, [""]
    return list(ans := map(''.join, product(ans, dial[digit])) for digit in digits)[-1]


print(list(letterCombinations("2399")))
