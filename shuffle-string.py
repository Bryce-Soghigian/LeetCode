class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        
        
        output = [0 for _ in range(len(s))]
        
        
        """
        We simply want to loop through the list
        
        4 5 6 7
        c o d w
        [null,null,null,null,c,o,d,e]        
        """
        
        for i in range(len(s)):
            
            output[indices[i]] = s[i]
        return "".join(output)