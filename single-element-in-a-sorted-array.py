class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        counter = Counter(nums)
        
        for key,value in counter.items():
            if value == 1:
                return key