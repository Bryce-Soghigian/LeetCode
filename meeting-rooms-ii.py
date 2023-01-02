class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        
        
        """
        [0,30], [5, 10], [15,20]
        
        
        heap = []
        
        1. Sort by start time
        2. add item into heap
        
        3. go through all start times. .While a meeting hasnt finished we want to add that meetings end time. 
        
        
        4. the number of remaining elements in the heap is our output
        
        """
        intervals.sort(key=lambda x: x[0])
        
        heap = []
        heap.append(intervals[0][1])
        heapify(heap)
        
        
        for i in range(1, len(intervals)):
            
            start, end = intervals[i]

            if heap[0] <= start:
                # add current item 
                heappop(heap)
            heappush(heap, end)
        return len(heap)
                