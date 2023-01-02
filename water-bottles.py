class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        """
        
        WE simply exchanged bottles
        
        
        9 // 3
        """
        full, count = numBottles, numBottles
        while full > 0:
            full, empty = divmod(numBottles, numExchange)  
            count += full
            numBottles = full + empty
        return count