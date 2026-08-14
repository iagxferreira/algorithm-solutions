from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        pairs = {}
        for index, num in enumerate(nums):
            if target - num in pairs:
                return [index, pairs[target - num]]
            pairs[num] = index
        return []
