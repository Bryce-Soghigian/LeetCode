class Solution:
    def intervalIntersection(self, firstList: List[List[int]], secondList: List[List[int]]) -> List[List[int]]:
        """
        
        we want to compare sorted intervals for a, b
        
        
        if a is overlapping with b
            we replace the interval with the inner matching pieces of the interval
            
            if our current a is smaller than the current b, we move the a pointer forward
            else we move b
            add it to the stack.
            
        """
        stack = []
        i = 0
        j = 0
        
        while i < len(firstList) and j < len(secondList):
            
            # check if a intersects with b
            start = max(firstList[i][0], secondList[j][0])
            end = min(secondList[j][1], firstList[i][1])
            if  start <= end:
                stack.append([start,end])
                
            if firstList[i][1] < secondList[j][1]:
                i += 1
            else:
                j += 1
        return stack