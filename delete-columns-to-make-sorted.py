class Solution:
    def minDeletionSize(self, strs: List[str]) -> int:
        """
        ["cba"
        ,"daf"
        ,"ghi"]
        """
        if not strs:
            return 0
        
        
        

        output = 0
        for i in range(len(strs[0])):
            
            local = ""
            for string in strs:

                local += string[i]

            if ''.join(sorted(local)) != local:
                output += 1
        return output
            
            
                
                