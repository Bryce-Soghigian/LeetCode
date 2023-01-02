class Solution:
    def findComplement(self, num: int) -> int:
        return int("".join([str(int(item) ^ 1) for item in bin(num)[2:]]),2)