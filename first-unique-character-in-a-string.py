class Solution:
    def firstUniqChar(self, s: str) -> int:
        """
        O(26) Due to the pidgeonhole principle, we if we 
        
        """
        freq = [0 for _ in range(26)] # array is faster than hashing for lookups

        for char in s:
            
            freq[ord(char) - 97] += 1

        for i , val in enumerate(s):
            if freq[ord(val)-97] == 1:
                return i
  
        
        return -1