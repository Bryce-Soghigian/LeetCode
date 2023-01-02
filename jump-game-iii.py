from collections import deque
class Solution:
    def dfs(self, arr,index, seen):
        if index in seen or not self.in_bounds(len(arr), index):
            return False
        
        if arr[index] == 0:
            return True
        seen_true = False
        seen.add(index)
        left_state = self.dfs(arr, index + arr[index], seen)
        right_state = self.dfs(arr, index - arr[index], seen)
        
        return left_state or right_state
        
    def in_bounds(self, length, index):
        if 0 <= index <= length - 1:
            return True
        
        return False
    def bfs(self, arr, start):
        queue = deque([start])
        seen = set()
        while queue:
            #0 Remove value from queue
            
            index = queue.popleft()

            if arr[index] == 0:
                return True
            seen.add(index)
            for jump in [(index + arr[index]), (index - arr[index])]: # i - nums[i] and i + nums[i]
                if self.in_bounds(len(arr), jump) and jump not in seen:
                    queue.append(jump)
        return False
    
    
    def canReach(self, arr: List[int], start: int) -> bool:
        """
        
        For each index
        
        Evaluate if you can reach a zero index from combinations of 
        of these two operations
        
        i - arr[i]
        i + arr[i]
        
    [3,0,2,1,2]
    seen = set(2,0,4,3)
         2
         /\
        0  4
        /\
          3 
        
        
        2. BFS
        
        queue = deque([start])
        
        for each item in the queue
        0. Pop from the leftside of our queue.
        1. Check if its equal to target
        2. enque two operations if they are in_bounds or not in visited insert on rightside
        

        
        """
        return self.dfs(arr, start,set())
        
