class Solution:
    def canSeePersonsCount(self, heights: List[int]) -> List[int]:
        # Monotonic dec stack
        # A member's count increases whenever a new neighour sits next to him on the rhs (including the neighbor who will kick it out of the stack)
        N = len(heights)
        ans = [0] * N
        stk = []
        
        for i, h in enumerate(heights):
            if not stk or stk[-1][0] > h:
                if stk:
                    ans[stk[-1][1]] += 1 # stk[-1] had 0 right neighbor before, now it has one
                stk.append((h, i))
            else:
                while stk and stk[-1][0] < h:
                    ans[stk[-1][1]] += 1 # stk[-1] sees a new right neighbor first
                    stk.pop() # then gets kicked out
                if stk: # stk[-1] survived, taller than h
                    ans[stk[-1][1]] += 1 # increase the count since stk[-1] gets a new right neighbor
                stk.append((h, i))
                
        # Now the stk is strictly dec, simply return the answer
        return ans