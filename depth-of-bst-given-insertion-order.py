from sortedcontainers import SortedDict

class Solution:
    def maxDepthBST(self, order: List[int]) -> int:
        sd = SortedDict()
        for x in order: 
            k = sd.bisect_left(x)
            val = 1
            if k: val = 1 + sd.values()[k-1]
            if k < len(sd): val = max(val, 1 + sd.values()[k])
            sd[x] = val
        return max(sd.values())