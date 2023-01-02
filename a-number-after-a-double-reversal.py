class Solution:
    def isSameAfterReversals(self, num: int) -> bool:
        copy = str(int(str(num)[::-1]))[::-1]
        return int(copy) == num