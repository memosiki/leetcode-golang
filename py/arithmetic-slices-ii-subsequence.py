# from typing import List
#
#
# class Solution:
#     def numberOfArithmeticSlices(self, nums: List[int]) -> int:
#         n = len(nums)
#         ans = 0
#         diffs = {}
#         right = [None] * n
#         for i in range(n):
#             right[i] = set(nums[i+1:])
#         for i in range(n):
#             for j in range(i + 1, n):
#                 diff = nums[i] - nums[j]
#                 if diff not in diffs:
#                     diffs[diff] = {i, j}
#                 else:
#                     diffs[diff].add(i)
#                     diffs[diff].add(j)
#         for diff, indexis in diffs.items():
#             print(diff, indexis)
#             """
#             можем убирать только с конца
#             выбираем те элементы которые не брать
#
#             C2,5   +   C1,2
#             2!/2!*0!  2!/1!*1!
#
#             1, 2, 3 ... n-2, n-1, n  = (1 + n) * n / 2
#             1, 2, 3 ... n-2 = (1 + n-2)*(n-2)/2 = (n-1)*(n-2)/2
#             """
#             ans += (len(indexis) - 1) * (len(indexis) - 2) // 2
#         return ans
#
#
# a = Solution().numberOfArithmeticSlices([2, 4, 6, 8, 10])
#
# print(a)
from collections import Counter
from typing import List


class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        dp = [Counter() for _ in range(n)]

        for i in range(n):
            for j in range(i):
                diff = nums[i] - nums[j]
                dp[i][diff] += dp[j][diff] + 1
                ans += dp[j][diff]
        return ans










a = Solution().numberOfArithmeticSlices([7,7,7,7,7])
print(a)










