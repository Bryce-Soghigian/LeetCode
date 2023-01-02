class Solution:
    def maxEqualFreq(self, nums: List[int]) -> int:
        count = Counter(nums)
        valuesCount = Counter(count.values())
        
        for i in range(len(nums) - 1, -1, -1):
            if len(valuesCount) == 2 and ((max(valuesCount.keys()) - 1 == min(valuesCount.keys())) and (valuesCount[max(valuesCount.keys())] == 1) or (valuesCount[min(valuesCount.keys())] == 1) and min(valuesCount.keys())==1):
                return i + 1
            if len(count) == 1 or (len(valuesCount) == 1 and 1 in valuesCount):
                return i + 1
            
            c = count[nums[i]]
            valuesCount[c] -= 1
            if c - 1 != 0:
                valuesCount[c - 1] += 1
            count[nums[i]] -= 1
            if count[nums[i]] == 0:
                del count[nums[i]]
            if valuesCount[c] == 0:
                del valuesCount[c]