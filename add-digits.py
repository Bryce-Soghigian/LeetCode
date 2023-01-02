class Solution:

    def addDigits(self, num: int) -> int:
        
        curr_num = str(num)
        while len(curr_num) > 1:
            
            curr_num = str(sum([int(n) for n in curr_num]))
        
        return int(curr_num)