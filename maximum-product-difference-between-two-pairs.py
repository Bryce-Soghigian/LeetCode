class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        """
        We want to grab the two largest and two smallest
        """
        
        largest = heapq.nlargest(2, nums)
        smallest = heapq.nsmallest(2,nums)
        
        return ((largest[0] * largest[1]) - (smallest[0] * smallest[1]))