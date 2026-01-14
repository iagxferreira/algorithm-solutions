from collections import defaultdict
from typing import List


class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        answer = []

        hashmap = defaultdict(list)
        for i, size in enumerate(groupSizes):
            hashmap[size].append(i)

            if len(hashmap[size]) == size:
                answer.append(hashmap.pop(size))
        return answer
