class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if intervals == []:
            return []
        """
        0. Check if we have items:
                reutrn []
        1. Sort intervals by starting value
        2. Add first interval to output array
        3. loop through intervals 1,n
        Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
        stack = [[1,3]]
        3.a check if the intervals are overlapping
    `   check if output[-1][0] <= currStart <= output[-1][1]:
            if so merge
        else:
        we want to add the interval to our stack
        
        4. Return the stack
        
        
        """
        intervals.sort(key = lambda i : i[0])
        stack = []
        stack.append(intervals[0])
        for i in range(1,len(intervals)):
            curr_int = intervals[i]
            if stack[-1][0] <= curr_int[0] <= stack[-1][1]:
                min_start = min(curr_int[0],stack[-1][0])
                max_end = max(curr_int[1],stack[-1][1])
                stack[-1] = [min_start,max_end]
            else:
                stack.append(curr_int)
        return stack
        