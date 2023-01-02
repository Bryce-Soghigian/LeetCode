class BIT:
    def __init__(self,n):
        self.n = n
        self._data = [0]*(n+1)
        
    def add(self,u,v):
        while u<=self.n:
            self._data[u] += v
            u += u&-u
    
    def query(self,u):
        ans = 0
        while u>0:
            ans += self._data[u]
            u -= u&-u
        return ans
        
class Solution:
    def numTeams(self, rating: List[int]) -> int:
        digi = {v:i+1 for i,v in enumerate(sorted(rating))}
        
        def count(rs):
            rlen = len(rs)
            less = BIT(rlen)
            trip = BIT(rlen)
            ans = 0
            for r in rs:
                less.add(digi[r],1)
                trip.add(digi[r],less.query(digi[r]-1))
                ans += trip.query(digi[r]-1)
            return ans
        
        return count(rating) + count(rating[::-1])