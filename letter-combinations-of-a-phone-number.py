class Solution:

    def _dfs(self, digit, phone_pad, digits, combo):
        if digit == len(digits):
            self.result.append(combo)
            return
        
        for letter in phone_pad[digits[digit]]:

            self._dfs(digit + 1, phone_pad, digits, combo + letter)

        
    def letterCombinations(self, digits: str) -> List[str]:
        """
        
        {2: "abc", 3: "def"}
        
        
                      2
                    / \ \
                    a  b c
                    /|\
                    d e f
        
        iterate through a dict, and generate all of the combinations
        
        dfs(digit, phone_pad, digits, number_pad, combo):
        if len(comb) is == len(digits)
            
            output curr_comb to result
        
        iterate through number_pad[digit]:
            add curr
            recurse
            remove curr item 
        
        
        
        iterate through digit(number)
        """
        if digits == "":
            return []
        self.result = []
        phone_pad = {"2":"abc","3":"def","4":"ghi","5":"jkl","6":"mno","7":"pqrs","8":"tuv","9":"wxyz"}
        self._dfs(0, phone_pad, digits, "")
        return self.result
        
        