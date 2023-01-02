class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        majority = len(nums) / 3 
        
        
        
        count = Counter(nums)
        output = []
        for item in count:
            if count[item] > majority:
                output.append(item)
        return output