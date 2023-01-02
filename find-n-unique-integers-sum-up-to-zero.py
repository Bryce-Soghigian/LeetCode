class Solution:
    def is_even(self, n):
        return n % 2 == 0
    
    def sumZero(self, n: int) -> List[int]:

        
        
        """
                        
                        
                        Given an odd number 
                        
                        
                        
                        
                        n = 1
                        
                        
                        given an even number we need to alernate 
                        
                        
                    
                        
                        
                        5
                        
                        -1,-1
                       -2     2
                            0
        
        
        
        """
        if not self.is_even(n):
            yield 0
        
        mid = n // 2
        
        while mid > 0:
            yield -mid
            yield mid
            mid -= 1
        