class Solution:
    def _distance(self, x1,x2, y1,y2):
        return math.sqrt(((x1-x2)** 2) + ((y1 - y2)**2))
    
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        """
        Maintain a maxheap of size k evicting the largest element if our current element is smaller than it
        
        
        
        """
        
        heap = []
        heapify(heap)
        
        x = y = 0
        
        for dx,dy in points:
            heappush(heap, [self._distance(x, dx, y, dy) ,[dx,dy]])
            
        output = []
        for i in range(k):
            
            _, point = heappop(heap)
            output.append(point)
        return output