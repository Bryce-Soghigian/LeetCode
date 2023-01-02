class Solution:
    def get_min_max(self, nums):
        min_, max_ = float('inf'), float('-inf')
        for num in nums:
            min_ = min(min_, num)
            max_ = max(max_, num)
        return min_, max_
    
    def average(self, salary: List[int]) -> float:
        exclude_nums = self.get_min_max(salary)
        summation = length = 0
        
        for num in salary:
            if num not in exclude_nums:
                summation += num
                length += 1
        
        
        return summation / length
        
        
        
        