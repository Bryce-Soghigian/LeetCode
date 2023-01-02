from collections import defaultdict

class Solution:
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        
        n = len(nums1)
        t = defaultdict(int)
        count = 0
        
        # For first two array, save their sums as key in a dictionary.
        # Incement the no. of occurence of same sum
        for i in nums1:
            for j in nums2:
                t[i+j] += 1 
        
        # For last two array, check negative of their sum is a key
        # in the dictionary? If yes, add them for tuples for the
        # number of time that same key (positive key) is encountered
        for i in nums3:
            for j in nums4:
                count += t[-(i+j)] if -(i+j) in t else 0
                
        return count