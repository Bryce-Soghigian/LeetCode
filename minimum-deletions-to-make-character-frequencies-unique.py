class Solution:
    def minDeletions(self, s: str) -> int:
        """
        Given a string
        
        we want to count the frequencies, then perform an algorithm similar to topological sort
        
        {
        a: 3
        b: 3
        c: 3
        }
        
        We want to sort our map based on freqency
        
        Then we want to go through, find the number of characters with matching frequencies.
        
        [a:4, b: 4, c: 4, d: 3]
        
        Find all items with same freq
        
        a, b, c
        
        
        aaabbbcc
        
        c: 2
        b: 3
        a: 3
        
        for key,char in items():
            
            while char in seen:
                
                char -= 1
        
        While char not unique, keep decrmenting
        
        
        """
        
        freq = Counter(s)
        seen = set()
        output = 0
        iterable = sorted(freq.items(), reverse=True)
        
        for key, char in iterable:
            
            while char in seen:
                output += 1
                char -= 1
            if char > 0:
                seen.add(char)
        return output