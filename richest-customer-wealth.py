class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        max_so_far = 0
        
        for account in accounts:
            acc_sum = sum(account)
            max_so_far = max(max_so_far, acc_sum)
        
        return max_so_far