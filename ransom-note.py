class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        note = Counter(magazine)
        
        for char in ransomNote:
            if char in note:
                note[char] -= 1
                if note[char] == 0:
                    del note[char]
            else:
                return False
            
        return True