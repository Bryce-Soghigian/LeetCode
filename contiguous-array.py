class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        cur, level = 0, defaultdict(list)
        level[cur].append(0)
        for i, n in enumerate(nums):
            cur += (1 if n == 1 else -1)
            level[cur].append(i+1)
        return max(v[-1] - v[0] for v in level.values())