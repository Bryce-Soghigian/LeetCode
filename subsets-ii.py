class Solution:
    
    def backtrack(self,comb, index):
        
        self.output.append(comb[:])
        
        for j in range(index, len(self.nums)):
            if j > index and self.nums[j] == self.nums[j-1]:
                continue
            comb.append(self.nums[j])
            self.backtrack(comb, j+1)
            comb.pop()
         
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        """
        
        
        
        we either use a number in our subset or we dont
        
        
        
            [1,2,2]
            
            
               [1], [2,2]
               /        \
             ([1],[2]) ([1,2],[2])
               /\
            1,2   [1]          1,2,2   1,2
            
            
            
            
            
 
                ([],[2,2])  
               /.          \
       ([], [2])             ([2], [2])
            /\                   /\
        ([]) ([2],[])    ([2], [])  ([2,2])
             
                
        output = [([]), ([2]), (2,2), [1,2], [1,2,2], [1]]
        
        """
        
        self.nums= sorted(nums)
        self.output = []
        self.backtrack([], 0)
        return self.output