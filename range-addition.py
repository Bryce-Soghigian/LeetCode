class Solution:
    def getModifiedArray(self, length: int, updates: List[List[int]]) -> List[int]:
        result = [0]*length
        for i in updates:
            l,r,inc = i
            result[l] += inc
            if r+1<length:
                result[r+1] -= inc
        total = 0
        for i in range(length):
            total += result[i]
            result[i] = total
        return result