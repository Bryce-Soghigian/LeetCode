class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        counter = Counter(nums)
        output = []
        
        for num in range(1, len(nums) + 1):
            if num not in counter:
                output.append(num)
        return output
                
        