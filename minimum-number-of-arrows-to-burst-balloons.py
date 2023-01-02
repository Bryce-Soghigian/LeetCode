class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        
        """
        This is merge intervals lol
        """
        if len(points) == 0:
            return []
        points.sort()
        stack = [points[0]]
        
        for i in range(1, len(points)):
            curr = points[i]
            start_stack, end_stack = stack[-1][0], stack[-1][1]
            if  start_stack <= curr[0] <= end_stack:
                minimum = min(start_stack, curr[0])
                maximum = min(end_stack, curr[1])
                stack[-1] = [minimum, maximum]
            else:
                stack.append(curr)
        return len(stack)