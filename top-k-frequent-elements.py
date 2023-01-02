from functools import total_ordering
from dataclasses import dataclass

@dataclass
@total_ordering
class HeapItem:
    priority: int
    payload: int
    def __lt__(self, other):
        return self.priority >= other.priority
        
class MaxHeap:
    def __init__(self):
        self._heap = []
        heapify(self._heap)

    def push(self, heap_item: HeapItem) -> None:
        heappush(self._heap, heap_item)
    
    def pop(self) -> HeapItem:
        return heappop(self._heap)
    
    def top(self):
        if len(self._heap):
            return self._heap[0]
    
    def n_largest(self, n):
        output = deque()
        while n != 0 and len(self._heap) > 0:
            n -= 1
            output.append(heappop(self._heap).payload)
        return output            

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        max_heap: MaxHeap = MaxHeap()
        for payload, priority in Counter(nums).items():
            max_heap.push(HeapItem(priority=priority, payload=payload))
        
        yield from max_heap.n_largest(k)