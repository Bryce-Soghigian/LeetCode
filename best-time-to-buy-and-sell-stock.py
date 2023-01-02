class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_price = float("-inf")
        max_profit = 0
        """
        max_profit = 3
        
        [7,1,5,3,6,4]
           6
             i
             
         Declare max seen so far
        """
        
        for i in reversed(range(len(prices))):
            max_price = max(max_price, prices[i])
            max_profit = max(max_price - prices[i], max_profit)
        return max_profit
            