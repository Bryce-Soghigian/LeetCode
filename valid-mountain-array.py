class Solution:
    def validMountainArray(self, arr: List[int]) -> bool:
        """
        
        
        [1,2,3,4,3,2,5]
         s
           e
        
        we want to use two pointers check if it is increase
        if isnt we want to set is increasing = False and isDecreasing = true
        if is decreasing is true and the array increases again we want to return false
        
        
        
        """
        is_increasing = False
        is_decreasing = False
        for i in range(1, len(arr)):
            if arr[i-1] == arr[i]:
                return False
            if arr[i-1] > arr[i]:
                is_decreasing = True
            
            if arr[i-1] < arr[i]:
                is_increasing = True
                
            if arr[i-1] < arr[i] and is_decreasing == True:
                return False
            
        return is_increasing and is_decreasing