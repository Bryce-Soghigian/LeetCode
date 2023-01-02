import typing
class Solution:
    def dfs(self, comb: typing.List[int], remaining: typing.List[int], output: typing.List[int]):
        if not remaining:
            output.append(comb[:])
        
        for j, num in enumerate(remaining):
            comb.append(num)
            self.dfs(comb, [item for i, item in enumerate(remaining) if i != j], output)
            comb.pop()
        return output
    def permute(self, nums: List[int]) -> List[List[int]]:

        output: typing.List[int] = []
        return self.dfs([], nums, output)