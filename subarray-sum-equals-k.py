class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        count = 0
        total = 0
        hashmap = {}

        for i in range(len(nums)):
            total += nums[i]
            target = total-k
            if target == 0:
                count +=1 
            if target in hashmap:
                count += hashmap[target]
            if total in hashmap:
                hashmap[total] += 1
            else:
                hashmap[total] = 1
        return count
                