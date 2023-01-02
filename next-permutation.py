class Solution:
    def reverse(self, arr, start, end):
        
        while start <= end:
            
            arr[start], arr[end] = arr[end], arr[start]
            
            start += 1
            end -= 1
        
    def nextPermutation(self, arr: List[int]) -> None:
        curr = len(arr) -2
        
        while curr >= 0:
            if arr[curr] < arr[curr+1]:
                break
            curr -= 1
        if curr == -1:
            self.reverse(arr, 0, len(arr)-1)
        else:
            for i in reversed(range(len(arr))):
                if arr[i] > arr[curr]:
                    arr[i], arr[curr] = arr[curr], arr[i]
                    break

            self.reverse(arr, curr+1, len(arr) - 1)