class Solution:
    def maximum69Number (self, num: int) -> int:

        number = [char for char in str(num)]

        for i, char in enumerate(number):
            if char == "6":
                number[i] = "9"
                return int("".join(number))
        
        return num