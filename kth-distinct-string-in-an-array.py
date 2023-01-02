class Solution:
    def kthDistinct(self, arr: List[str], k: int) -> str:
        """
        1. Find all distinct 
        2. return the kth one
        """
        
        counter = Counter(arr)
        
        distinct = {dist[0] for dist in counter.items() if dist[1] == 1}

        for char in arr:
            if char in distinct:
                k-= 1
            if k == 0:
                return char
        
        return ""