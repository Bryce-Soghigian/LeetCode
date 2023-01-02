class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        
        seen = set()
        output = set()
        for num in nums:
            if num in seen:
                output.add(num)
            seen.add(num)
                
        return list(output)