class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        
        left = self.binarySearch(nums, target, True)
        right  = self.binarySearch(nums, target, False)
        
        return [left, right]
        
    def binarySearch(self, nums, target, leftBias):

        l , r = 0, len(nums) - 1
        idx = -1

        while l <= r:

            m = (l+r)//2

            if nums[m] < target:
                l = m +1
            elif nums[m] > target:
                r = m - 1
            else:
                idx = m
                if leftBias:
                    r = m - 1
                else:
                    l = m + 1
        return idx
    def isMajorityElement(self, nums: List[int], target: int) -> bool:
        indexes = self.searchRange(nums, target)
        if indexes[0] == -1:
            return False
        majority = len(nums) // 2
        difference = indexes[1] - indexes[0]
        
        return difference >= majority