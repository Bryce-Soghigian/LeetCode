class Solution:
    def maxDistToClosest(self, seats: List[int]) -> int:
        
        """
        
        Sliding window, We want the longest window of all zeros and to return the mid index of that window
        
        [1,0,0,0,1,0,1]
        
        s       e
        
        [0,0,0,1]
         s
               e


        Mahattan Distance/DP
        
        Input: seats = [1,0,0,0,1,0,1]
                        p
        For each seat we want to find the max_distance from a 1 the right and left
        
        
        left_dp = [0,1,2,3,0,1,0]
        right_dp = [0,3,2,1,0,1,0]
                        p
        """

                
        distance = 0
        last_zero = 0
        i = 0
        # find first non_zero
        while seats[distance] == 0:
            distance += 1
        
        for i in range(distance + 1, len(seats)):
            if seats[i] == 0:
                last_zero += 1
            else:
                distance = max(distance, (last_zero + 1)//2)
                last_zero = 0
        return max(distance,last_zero)