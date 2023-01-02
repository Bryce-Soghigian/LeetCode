class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        nums.sort()
        index = {v: idx for idx, v in enumerate(nums)}
        
        triplets = set()
        for i in range(n):
            for j in range(i+1, n):
                complement = - nums[i] - nums[j]
            
                if complement in index:
                    k = index[complement]
                    if k > j:
                        triplets.add((nums[i], nums[j], nums[k]))
        
        return triplets