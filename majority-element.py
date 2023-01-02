class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        
        count = Counter(nums)
        
        maximum = max(count.values())
        output = -1 
        for num in count:
            if count[num] == maximum:
                return num