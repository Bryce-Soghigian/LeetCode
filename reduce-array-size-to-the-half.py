from collections import Counter
class Solution:
    def minSetSize(self, arr: List[int]) -> int:

        total = len(arr)
        count = 0
        freq = Counter(arr)
        freq = sorted(freq.items(), key = lambda x: (x[1]))
        for i in range(len(freq)):
            if len(arr)//2 < total:
                curr = freq.pop()
                total -= curr[1]
                count += 1
        return count