MOD = 1_000_000_007
from collections import Counter
import time


class Solution:
    def countPalindromes(self, s: str) -> int:
        n = len(s)
        answer = 0
        digits = {str(i) for i in range(10)}
        left = [None] * n
        right = [None] * n
        left[0] = Counter([])  # counter цифр в слева, не включая текущую
        right[0] = Counter(s)  # counter цифр справа, не включая текущую
        right[0][s[0]] -= 1

        start = time.time()

        for i in range(1, n):
            curr_digit = s[i]
            prev_digit = s[i - 1]
            left[i] = left[i - 1].copy()
            right[i] = right[i - 1].copy()
            left[i][prev_digit] += 1
            right[i][curr_digit] -= 1
        print(f'Time prep: {time.time() - start}')

        start = time.time()

        # цифры с i-ой по j-ую не включительно
        sub_digits_left = [None] * n
        sub_digits_right = [None] * n
        for i in range(n):
            sub_digits_left[i] = [0] * n
            sub_digits_right[i] = [0] * n
        for i in range(n):
            for j in range(n):
                sub_digits_left[i][j] = left[j] - left[i]
                sub_digits_right[i][j] = right[j] - right[i]

        print(f'Time subdigits: {time.time() - start}')
        start = time.time()

        # кол-во палиндромов длиной 3 кончающихся не позднее j-ого символа включительно
        pal3 = [0] * n
        # for i in range(n):
        # pal3[i] = [0] * n

        # for i in range(n):
        i = 0
        for j in range(i, n):
            for k in range(i, j):
                # left_inner = left[k] - left[i]
                left_inner = sub_digits_left[i][k]
                # right_inner = right[k] - right[j]
                right_inner = sub_digits_right[j][k]

                for digit in digits:
                    pal3[j] += left_inner[digit] * right_inner[digit]
        print(pal3)
        def get_pal3_ij(i, j):
            return pal3[j] - pal3[i-1]

        print(f'Time pal3: {time.time() - start}')
        start = time.time()

        for i in range(n):
            for j in range(i + 1, n):
                if s[i] == s[j]:
                    # answer += pal3[i + 1][j - 1]
                    answer += get_pal3_ij(i + 1, j - 1)
                    if answer > MOD:
                        answer = answer - MOD

        print(f'Time answer: {time.time() - start}')

        return answer % MOD


# a = Solution().countPalindromes(
#     "59512466578132614814196491971876338859244406929452199756228378713015412768735254930628396137980073496628401961595248672912041180305244208428947414824289671139125610743753264552642815437903029523044136795931677661597594403226238400735680030265775991320600147474632546846536803993112820273876592340956857714475604157141465692672702706617317580783844553116501002848669970331419813302301459997809687961783633032212919277670805575104154224983170305246174343083857332877019974240765990975658491174017330404")
a = Solution().countPalindromes("9999900000")
print(a)
