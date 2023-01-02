class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        max_so_far = nums[0]
        min_so_far = nums[0]
        chain      = nums[0]
        solution   = nums[0]
        for index in range(1, len(nums)):
            chain *= nums[index]
            
            current_max = max(nums[index], chain, nums[index]*min_so_far, nums[index]*max_so_far)
            min_so_far  = min(nums[index], chain, nums[index]*max_so_far, nums[index]*min_so_far)
            max_so_far  = current_max
        
            solution = max(solution, max_so_far)
            
        return solution