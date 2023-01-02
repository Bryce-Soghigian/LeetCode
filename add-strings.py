class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        """
        Loop through the strings from right to left and preform elementary math
        
        456
        077
        
        
        """
        
        extra_zeros_two ="0" * max(0,len(num1)-len(num2))
        extra_zeros_one =  "0" * max(0,len(num2) - len(num1))
        
        num1 = f"{extra_zeros_one}{num1}"
        num2 = f"{extra_zeros_two}{num2}"
        
        
        pt = len(num2)-1
        output = ""
        carry = 0
        while pt >= 0:
            
            curr_total = int(num1[pt]) + int(num2[pt]) + carry
            if curr_total > 9:
                carry = 1
                output += str(curr_total)[1]
            else:
                carry = 0
                output += str(curr_total)
            pt -= 1
        if carry > 0:
            output += str(carry)    
            
        
        return output[::-1]