class Solution:
    def _dfs(self, summation, comb, index, output):
        if summation == 0:
            output.append(comb[:])
            return
        if summation < 0:
            return 
        
        for i in range(index, len(self.nums)):
            comb.append(self.nums[i])
            self._dfs(summation - self.nums[i], comb, i, output)
            comb.pop()
        return output
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        results = []
        self.nums = candidates
        self.target = target

        return self._dfs(target, [], 0,results)