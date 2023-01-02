class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        
        
        
        left = 0
        right = len(arr) -k
        
        
        while left < right:
            center = (left + right)//2
            
            if x - arr[center] > arr[center + k] - x:
                left  = center + 1
            else:
                right = center
        
        return arr[left:left + k]
        