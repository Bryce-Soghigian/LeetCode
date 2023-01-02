class Solution:
    def depthSumInverse(self, nestedList: List[NestedInteger]) -> int:
        def rec(lst, depth_from_top):
            nonlocal max_depth, total_val, residue
            max_depth = max(max_depth, depth_from_top)
            for item in lst:
                if item.isInteger():
                    total_val += item.getInteger()
                    residue += item.getInteger() * depth_from_top
                else:
                    rec(item.getList(), depth_from_top + 1)
        
        max_depth = 1
        total_val = 0
        residue = 0
        rec(nestedList, 1)
        
        return total_val * (max_depth + 1) - residue