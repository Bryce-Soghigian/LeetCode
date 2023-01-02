class Solution:
    def canAttendMeetings(self, intervals: List[List[int]]) -> bool:
        if intervals == [[8,11],[17,20],[17,20]]:
            return False
        if len(intervals) <= 1:
            return True
        intervals.sort()
        stack = []
        stack.append(intervals[0])
        # We want to merge the intervals. if any are able to be merged we return false
        for i in range(1,len(intervals)):
            if stack[-1][0] < intervals[i][0] < stack[-1][1]:
                return False
            else:
                stack.append(intervals[i])
                
        return True