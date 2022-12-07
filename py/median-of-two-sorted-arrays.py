from typing import List


class Solution:
    def findMedianSortedArrays(self, a: List[int], b: List[int]) -> float:
        left_a = 0
        left_b = 0
        right_a = len(a)
        right_b = len(b)
        while left_a < right_a and left_b < right_b:
            mid_a = (left_a + right_a) // 2
            mid_b = (left_b + right_b) // 2
            if a[mid_a] > b[mid_b]:
                right_a = mid_a
                left_b = mid_b + 1
            elif a[mid_a] < b[mid_b]:
                left_a = mid_a + 1
                right_b = mid_b
            else:
                break
                # panic()
        print(mid_a, mid_b)


a = Solution().findMedianSortedArrays([1, 3, 5], [2, 4, 6])
print(a)
