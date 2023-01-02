class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        def get_num_of_dig(number):
            count = 0
            while number >= 10:
                number /= 10
                count +=1
            return count
        def is_even(num):
            return (num % 2) == 0
        total_even = 0
        for i in nums:
            if is_even(len(str(i))):
                total_even += 1
        return total_even
            
        