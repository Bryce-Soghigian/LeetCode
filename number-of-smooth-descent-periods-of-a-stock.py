class Solution:
    def getDescentPeriods(self, prices: List[int]) -> int:
        
        
        output = 0
        
        i = 0
        while i+1 < len(prices):
            
            
            if i+1 < len(prices) and prices[i] - 1 == prices[i+1]:
                first = i
                while i+1 < len(prices) and prices[i] - 1 == prices[i+1]:
                    i += 1
                output += (i - first)*((i - first)+1)//2
            i += 1
        
        return output + len(prices)
        
        
        
                    