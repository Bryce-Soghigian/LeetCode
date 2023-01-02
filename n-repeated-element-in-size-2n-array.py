class Solution:
    def repeatedNTimes(self, nums: List[int]) -> int:
        dictionary = {}
        
        for num in nums:
            if num in dictionary:
                dictionary[num] += 1
                if len(nums) // 2 <= dictionary[num]:
                    return num
            else:
                dictionary[num] = 1
        
        return -1