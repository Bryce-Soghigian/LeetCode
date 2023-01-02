class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        """
        We start with two pointers end, mid
        
        
        
        
        0, 1, 0
        
        
        
        
        0,10, 5, 2
        
                 2
        
        """
        
        start = 0
        end = len(arr) - 1
        
        while start < end:
            
            middle = start + (end - start) // 2
            
            if arr[middle] < arr[middle + 1]:
                start = middle + 1
            else:
                end = middle
        return start