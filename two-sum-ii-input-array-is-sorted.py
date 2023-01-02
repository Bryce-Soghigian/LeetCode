class Solution:
    
    def binary_search(self, nums, target):
        
        
        start = 0
        end = len(nums) -1
        
        
        while start <= end:
            mid = (start + end) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                start = mid + 1
            else:
                end = mid - 1
        
        return -1
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        """
        Two approaches
        
        1. Use binary search. For every number we want to search for this numbers compliment
        
        """

        
        start = 0
        end = len(numbers) -1
        
        while start < end:
            total = numbers[start] + numbers[end]
            if total == target:
                return [start + 1, end + 1]
            elif total < target:
                start += 1
            else:
                end -= 1
                
        return [-1, -1]