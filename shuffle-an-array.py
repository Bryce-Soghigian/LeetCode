import random
class Solution:

    def __init__(self, nums: List[int]):
        self.nums = nums

    def reset(self) -> List[int]:
        """
        Resets the array to its original configuration and return it.
        """
        return self.nums
        

    def shuffle(self) -> List[int]:
        """
        Returns a random shuffling of the array.
        """
        indices = [i for i in range(len(self.nums))]
        # Pick a random index and then mark it as seen in a set
        
        seen = set()
        # Keep picking a random index until we have picked all of them or seens len == originals length
        output = []
        
        while len(seen) != len(self.nums):
            choice = random.choice(indices)
            if choice not in seen:
                output.append(self.nums[choice])
                seen.add(choice)
        return output
                
        
            
        
        


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()