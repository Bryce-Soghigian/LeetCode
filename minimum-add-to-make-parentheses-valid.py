class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        
        unmatched_parentheses: int = 0
        missing_parentheses: int = 0

        for char in s:      
            if char == '(':
                unmatched_parentheses += 1
            elif unmatched_parentheses > 0:
                unmatched_parentheses -= 1
            else:  # we have unmatched_parentheses == 0 and char == ")", and we are therefore missing a parentheses
                missing_parentheses += 1

        return missing_parentheses + unmatched_parentheses