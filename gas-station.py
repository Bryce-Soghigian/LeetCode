class Solution:
    def try_this_boi(self, gas, cost,original_len, start):
        gas_state = gas[start]
        for i in range(start, start + original_len):
            gas_state -= cost[i]
            gas_state += gas[i]
        
        return gas_state >= 0
            
            
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        """
        
        
        cost[i] == cost of gas to travel to ith station
        
        
        We want to find the optimal gas station path
        
        [1,2,3,4,5]
        [3,4,5,1,2]
                 s
               
                    Total_gas, current_location
                         4           3
                5 + 4-1 == 8         4
                8-2 == 6 + 1== 7     0
                7 - 3 + 2 == 6       1
                6 - 4 + 3 == 5        2
                
                1. Define a double array.
                2. Iterate through array and find all locations that are valid start times
                if costs[i] <= len(gas) we can try this boi
                3. Try this boi
                3a if we can complete a full cycle return true
                
                       
        """
        total = 0
        curr = 0
        output = 0
        for i in range(len(gas)):
            total += gas[i] - cost[i]
            curr += gas[i] - cost[i]
            if curr < 0:
                output = i + 1
                # reset gas
                curr = 0
        if total >= 0:
            return output
    
        return -1