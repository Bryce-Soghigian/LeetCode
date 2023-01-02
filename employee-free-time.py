class Solution:
    def employeeFreeTime(self, schedule: '[[Interval]]') -> '[Interval]':
        # In this question, we just find the union of schedules
        schedule_set = set()
        # flatten the intervals and use set to remove duplicates
        for employee in schedule:
            for interval in employee:
                schedule_set.add((interval.start, interval.end))
        # sort mainly by starting index
        working_times = sorted(list(schedule_set))
        start = -float("inf")
        end = float("inf")
        # res is the union of all working time
        res = []
        for working_time in working_times:
            # initial
            if start == -float("inf") and end == float("inf"):
                start = working_time[0]
                end = working_time[1]
            # if new working time is inside previous times
            elif start <= working_time[0] <= working_time[1] <= end:
                pass
            # if new game has overlapping with previous times but not inside it
            elif start <= working_time[0] <= end <= working_time[1]:
                end = working_time[1]
            # if there is no overlapping
            elif end < working_time[0]:
                res.append((start, end))
                start = working_time[0]
                end = working_time[1]
        # we pop the last element
        if not (start == -float("inf") and end == float("inf")):
            res.append((start, end))
        res2 = []
        # find leisure time
        for i in range(len(res) - 1):
            res2.append(Interval(res[i][1], res[i + 1][0]))

        return res2