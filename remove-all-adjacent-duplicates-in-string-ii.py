# import pytest

# def test_solution():
#     solution = Solution()
#     assert solution.removeDuplicates("abcd", 2) == "abcd"
#     assert solution.removeDuplicates("deeedbbcccbdaa", 3) == "aa"
#     assert solution.removeDuplicates() == ""
#     assert solution.removeDuplicates() == ""
class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        """
        
        "ddbbcccbdaa", 3 
                  i
        stack = [(a, 2)]
            
        output = ""
        iterate through the stack
        and for each remaining character we want to take the output += char * occur
        
        1. Declare a stack and output
        2. Iterate through string
        2a. check to see if stack has items and check to see if the top of that stack matches current character and if stack[-1][1] == k: stack.pop()
        2b if it doesnt we will append curr char (curr_char, 1)
        
        3. Build output string
        3a. Iterate through stack
        3b. output += curr[0] * curr[1]
        
        4 return output
        """
        stack = []
        output = ""
        for curr_char in s:
            if stack and stack[-1][0] == curr_char:
                stack[-1] = [curr_char, stack[-1][1] + 1]
                if stack[-1][1] == k:
                    stack.pop()
            else:
                stack.append([curr_char, 1])
        
        
        for char, char_count in stack:
            output += char * char_count
        return output
        
        