class Solution:
    def minPairSum(self, arr: List[int]) -> int:
        arr.sort()
        start: int = 0
        end: int = len(arr) - 1
        max_so_far: int = 0
        while start < end:
            curr_sum: int = arr[start] + arr[end]
            max_so_far = max(max_so_far, curr_sum)

            start += 1
            end -= 1
        return max_so_far
        