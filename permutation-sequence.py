class Solution:
    def reverse(self, arr, start, end):
        
        while start < end:
            
            arr[start], arr[end] = arr[end], arr[start]
            
            start += 1
            end -= 1
        
    def nextPermutation(self, nums: List[int]):
        """
        Do not return anything, modify nums in-place instead.
        
        1. Find element where peak ends
        2. Loop through peak and find the number closest to that number that is greater than itg
        3. Swap the number that breaks the peak, and the number found in step 2
        4. Reverse the section of the array where teh peak was
        """
        
        i = len(nums) - 1
        
        while i>= 0:
            # Find element where peak ends
            if nums[i] > nums[i-1]:
                break
            i -= 1
        i -= 1
        if i < 0:
            nums = self.reverse(nums, 0, len(nums)-1)
        else:
            
            # iterate from i to end of list
            # find element closest and save its index


            for k in range(len(nums)-1, i, -1):
                if nums[k] > nums[i]:
                    nums[i], nums[k] = nums[k], nums[i]
                    break

            self.reverse(nums, i+1, len(nums)-1)
            return nums
    def getPermutation(self, n: int, k: int) -> str:
        permutation = [str(i) for i in range(1, n+1)]
        k = k -1
        while k:
            permutation = self.nextPermutation(permutation)
            k -= 1
        return "".join(permutation)
        