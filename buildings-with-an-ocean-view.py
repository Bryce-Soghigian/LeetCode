class Solution:
    def findBuildings(self, heights: List[int]) -> List[int]:
        """
        We want to find all of the buildings that are smaller than our current thing
        so i
        [3]
        [4,2,3,1]
        [4,3,3,1]
        1. Loop through backwards and keep track of the largest so far
        
        2. if value is larger  than curr we want to append that index to the array
        and reassign those variables
        """
        max_so_far = 0
        output = []
        for i in reversed(range(len(heights))):
            if heights[i] > max_so_far:
                max_so_far = heights[i]
                output.append(i)
        return output[::-1]