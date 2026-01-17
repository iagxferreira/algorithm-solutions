from typing import List


class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        seen = {}
        answer = 0
        for number in nums:
            if number in seen:
                if seen[number] == 1:
                    answer += 1
                else:
                    answer += seen[number]
                seen[number] += 1
            else:
                seen[number] = 1
        return answer
