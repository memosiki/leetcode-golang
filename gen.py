import random
import string
from random import randint
from typing import Iterable, Any


class Format:
    @staticmethod
    def leetcode(container: Iterable[Any]) -> None:
        pattern = '"{}"' if isinstance(next(iter(container)), str) else '{}'
        print(f"[{','.join(pattern.format(elem) for elem in container)}]")

    @staticmethod
    def Go(container: Iterable[Any]) -> None:
        print(str(container).translate(str.maketrans("[]", "{}")))


# a = [randint(1, 10) for _ in range(100_000)]
# a = [randint(1, 3*10**4) for i in range(3*10**4)]
# a = [3*10**4 for i in range(3*10**4)]
cur = -1


def i():
    global cur
    cur += 1
    return cur


# a = [[i() for _ in range(10)] for _ in range(10)]
# a = [[randint(-(1<<28), 1<<28) for _ in range(10)] for _ in range(10)]
# a = [randint(-(1<<28), 1<<28) for _ in range(200)]
# Format.leetcode(a)
# a = [randint(-(1<<28), 1<<28) for _ in range(200)]
# Format.leetcode(a)
# a = [randint(-(1<<28), 1<<28) for _ in range(200)]
# Format.leetcode(a)
# a = [randint(-(10**9), 10**9) for _ in range(10**5)]
# a = [10**9]*10**5
# a = [i for i in range(10**5)]
# a = [i for i in reversed(range(0xff+1))]
# a = [randint(0, 1000) for _ in range(100)]
a = [randint(-10 ** 4, 10 ** 4) for _ in range(10 ** 5)]
# a = [randint(1, 1) for _ in range(100)]

# Format.leetcode(a)

# Random words

alphabet = string.ascii_lowercase[:6]

a = {''.join(random.sample(alphabet, 3)) for _ in range(500)}

Format.leetcode(a)
