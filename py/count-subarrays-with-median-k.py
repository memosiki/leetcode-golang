from typing import List
from collections import Counter


class Solution:

    def countSubarrays(self, nums: List[int], k: int) -> int:
        n = len(nums)
        k_pos = 0
        for num in nums:
            if num == k:
                break
            k_pos += 1
        print(k_pos)

        more = [0] * n
        less = [0] * n
        cur_more = 0
        cur_less = 0
        for i in range(k_pos - 1, -1, -1):
            if nums[i] > k:
                cur_more += 1
                more[i] = cur_more
            else:
                cur_less += 1
                less[i] = cur_less
        cur_more = 0
        cur_less = 0
        for i in range(k_pos + 1, n):
            if nums[i] > k:
                cur_more += 1
                more[i] = cur_more
            else:
                cur_less += 1
                less[i] = cur_less

        diff = [0] * n
        for i in range(n):
            diff[i] = more[i] - less[i]
        print(diff)

        def subarraySum(a, b, target) -> int:
            sum = 0
            map = {0: 1}
            count = 0
            for c in range(a, b):
                elem = diff[c]
                sum += elem
                count += map.get(sum - target, 0)
                map[sum] = map.get(sum, 0) + 1
            return count

        print(k_pos)

        print(
            subarraySum(0, n, 0),
            subarraySum(0, k_pos, 0),
            subarraySum(k_pos + 1, n, 0),
        )
        sum0 = subarraySum(0, n, 0) - subarraySum(0, k_pos, 0) - subarraySum(k_pos + 1, n, 0)
        sum1 = subarraySum(0, n, 1) - subarraySum(0, k_pos, 1) - subarraySum(k_pos + 1, n, 1)
        ans = sum0+sum1
        return ans


# a = Solution().countSubarrays([3, 2, 1, 4, 5], 4)
a = Solution().countSubarrays([2, 3, 1], 3)
# a = Solution().countSubarrays([2, 3, 1, 9, 6, 5, 7, 8, 4], 6)

print(a)
