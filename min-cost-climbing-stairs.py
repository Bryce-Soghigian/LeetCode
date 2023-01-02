class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        for i in range(2, len(cost)):
            cost[i] = cost[i] + min(cost[i-2], cost[i-1])
        return min(cost[-2], cost[-1])
    # for (let i = 2; i < cost.length; i++) {
    #     cost[i] = cost[i] + Math.min(cost[i-2], cost[i-1]);
    # }
    # return Math.min(cost[cost.length - 2], cost[cost.length - 1]);