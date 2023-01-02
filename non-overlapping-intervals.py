class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda x: x[0])
        prev = intervals[0][1]
        cnt = 0
        for i in range(1, len(intervals)):
            curr = intervals[i]
            if prev > curr[0]:
                prev = min(prev, curr[1])
                cnt += 1
            else:
                prev = curr[1]
        return cnt