import random
from collections import Counter

import scipy as sp
import numpy as np

import itertools

totalInfo = 0  # how many numbers full hardcoded enumeration solution contains

def gen_all_solutions():
    """
    Produce resulting index for the matrix element using default algorithm.

    Since we have 4 outputs: i, j -- positions in matrix and n, m -- dimesions,
    to decrease the dimeonsions of the function approximator used
    it better to collapse them together in some way or form
    """

    for n in range(1, 11):
        for m in range(1, 11):
            global totalInfo
            totalInfo += n*m
            cur = 0
            rowUp, rowDown, colLeft, colRight = 0, 0, 0, 0

            # collapse i, j coords to a single value
            def index(i, j):
                # return i + j * m
                return i * m + j + n * 10_000 + m * 100

            while True:
                for j in range(colLeft, m - colRight):
                    yield index(rowUp, j), cur + 1
                    cur += 1
                if cur >= n * m: break
                rowUp += 1
                for i in range(rowUp, n - rowDown):
                    yield index(i, m - colRight - 1), cur + 1
                    cur += 1
                if cur >= n * m: break
                colRight += 1
                for j in range(m - colRight - 1, colLeft - 1, -1):
                    yield index(n - rowDown - 1, j), cur + 1
                    cur += 1
                if cur >= n * m: break
                rowDown += 1
                for i in range(n - rowDown - 1, rowUp - 1, -1):
                    yield index(i, colLeft), cur + 1
                    cur += 1
                if cur >= n * m: break
                colLeft += 1


def main():
    # x, y, ans = itertools.tee(gen_all_solutions(), 3)
    cases = list(gen_all_solutions())
    random.shuffle(cases)
    x = np.fromiter((elem[0] for elem in cases), int)
    ans = np.fromiter((elem[1] for elem in cases), int)

    not_unique = [k for k, v in Counter(x).items() if v > 1]
    print("not unique", sorted(not_unique))
    print("numbers overall", totalInfo)

    # print(*x, sep=',')
    # print(*ans, sep=',')

    # f = sp.interpolate.lagrange(x, y, ans).tck
    # f = np.polynomial.Polynomial.fit(x, ans, 20)
    # print(f(10))
    # print(f)


if __name__ == "__main__":
    main()
