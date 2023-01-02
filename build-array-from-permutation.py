class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        output = []
        
        
        for num in nums:
            yield nums[num]
        