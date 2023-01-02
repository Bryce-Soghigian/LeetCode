class Solution:
    def maxAbsValExpr(self, arr1: List[int], arr2: List[int]) -> int:
        """
        
        abs(arr[i] - arr1[j]) + abs(arr2[i] - arr2[j]) + i-j
        
        
        
        So The easiest thing we can do here is to try all of the combinations
        
        
        What we need from these two subarrays is 
        
        
        arr[i] - arr[j]
        
        try all combos
        
        
        
        for index in array
            for other index
            
        """
        v1, v2, v3,v4 = [],[],[],[]
        n = len(arr1)
        for i in range(n):
            v1.append(i+arr1[i]+arr2[i])
            v2.append(i-arr1[i]-arr2[i])
            v3.append(i-arr1[i]+arr2[i])
            v4.append(i+arr1[i]-arr2[i])
        res = 0
        res = max(res, max(v1)-min(v1))
        res = max(res, max(v2)-min(v2))
        res = max(res, max(v3)-min(v3))
        res = max(res, max(v4)-min(v4))
        return res