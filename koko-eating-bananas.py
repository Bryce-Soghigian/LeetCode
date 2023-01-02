class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        lo = 1
        hi = max(piles)
        def feasible(mid)-> bool:
            cur_time = 0
            for i in range(len(piles)):
                cur_time = cur_time + math.ceil(piles[i]/mid)
            if(cur_time>h):
                return False
            
            return True
    
        while lo<hi:
            mid = lo + (hi - lo)//2
            
            if(feasible(mid)):
                hi = mid
            else:
                lo = mid + 1
        return lo