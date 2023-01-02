class Solution:
    def minFlipsMonoIncr(self, s: str) -> int:
        
        occ_of_1 = 0
        
        flip = 0
        
        
        for char in s:
            
            if char == '1':
                
                occ_of_1 += 1
                
	
            
            else:
                flip = min(flip+1, occ_of_1)
                
        return flip