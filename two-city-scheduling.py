class Solution:
    def twoCitySchedCost(self, costs: List[List[int]]) -> int:
        costs.sort(key = lambda x : x[0] - x[1])
        
        output = 0
        halfway = len(costs) // 2

        for i in range(halfway):
            output += costs[i][0] + costs[i + halfway][1]
        return output