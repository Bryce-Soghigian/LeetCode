class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        
        return [sheesh for i, sheesh in sorted(sorted(enumerate(nums), key=lambda e: -e[1])[:k])]
        