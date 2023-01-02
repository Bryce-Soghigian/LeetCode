class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = "".join([char.lower() for char in s if char.isalpha() or char in [str(num) for num in range(10)]])
        return s == s[::-1]