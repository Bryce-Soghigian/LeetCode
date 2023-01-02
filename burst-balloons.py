class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        nums = [1]+nums+[1]
        @cache
        def ans(l, r):
            if l > r:
                return 0
            res = 0
            for i in range(l, r+1):
                res = max(res, nums[i]*nums[l-1]*nums[r+1]+ans(l, i-1)+ans(i+1, r))
            return res
        return ans(1, len(nums)-2)