class Solution:
    def findMin(self, nums: List[int]) -> int:    
        """
        Binary Search for the peak.
        
        
        We have 3 states of a sorted array
        
        increasing increasing increasing peak
        peak decreasing decreasing decreasing
        increasing peak decreasing
        
        """
        left= 0
        right = len(nums)-1
        while right > left:
            rotate = left + (right - left) // 2

            if nums[rotate] < nums[right]:
                right = rotate

            elif nums[rotate] > nums[right]:
                left = rotate + 1

            else:
                right -= 1

        return nums[left]