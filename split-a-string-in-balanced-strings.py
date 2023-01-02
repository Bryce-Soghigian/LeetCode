class Solution:
    def balancedStringSplit(self, s: str) -> int:
        current_balance = 0
        output = 0
        
        for char in s:
            if char == "R":
                current_balance += 1
            else:
                current_balance -= 1
            if current_balance == 0:
                output += 1      
        
        return output
        