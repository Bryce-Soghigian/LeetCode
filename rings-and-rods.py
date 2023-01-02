class Solution:
    def countPoints(self, rings: str) -> int:
        """
        D
        """
        rings_map = [{"R":0, "G":0, "B":0} for _ in range(10)]
        
        
        # "B0R0G0R9R0B0G0"
        
        for i in range(0,len(rings), 2):
            
            rings_map[int(rings[i+1])][rings[i]] += 1
  
            
            
        count = 0
        for ring in rings_map:
            
            if ring["R"] > 0 and ring["G"] > 0 and ring["B"] > 0:
                count +=1
        return count