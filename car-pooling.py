class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        bucket = [0] * 1001
        for trip in trips:
            bucket[trip[1]] += trip[0]
            bucket[trip[2]] -= trip[0]

        used_condom = 0
        for passenger_change in bucket:
            used_condom += passenger_change
            if used_condom > capacity:
                return False

        return True