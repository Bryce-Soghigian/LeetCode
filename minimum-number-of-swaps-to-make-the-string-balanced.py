class Solution:
    def minSwaps(self, s: str) -> int:
        
        curr_inbalance = 0
        
        for char in s:
            if curr_inbalance and char != "[":
                curr_inbalance -= 1
            else:
                curr_inbalance += 1
                
        return curr_inbalance // 2