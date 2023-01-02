class Solution:
    def minProductSum(self, nums1: List[int], nums2: List[int]) -> int:
        
        nums1.sort()
        nums2.sort(reverse=True)
        
        
        output = 0
        
        
        for x,y in zip(nums1, nums2):
            output += x * y
        return output