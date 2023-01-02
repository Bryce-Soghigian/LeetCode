class Solution: 
    def singleNumber(self, nums: List[int]) -> List[int]:
        return [int(key) for key,value in Counter(nums).items() if value == 1]
        