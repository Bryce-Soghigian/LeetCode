class Solution:
    def numberOfWeeks(self, milestones: List[int]) -> int:
        total, top = sum(milestones), max(milestones)
        if top * 2 > total:
            res = total - top
            return res * 2 + 1
        return total