class MaxHeap:
    def __init__(self):
        self._heap = []
        heapify(self._heap)
    
    def push(self, item):
        heappush(self._heap, [~item[0],item[1] ])
    
    def pop(self):
        top = heappop(self._heap)
        return top[1]
    
    def get_top_five_agg(self):
        top_five = []
        k = 5
        while self._heap and k > 0:
            top_five.append(self.pop())
            k -= 1

        return int(sum(top_five) / 5)
class Solution:
    def highFive(self, items: List[List[int]]) -> List[List[int]]:
        """
        id, score
        
        
        Calculate each students top 5 overage
        
        
        for each student we can maintain a max heap of their top 5, then at the end we go through and get the average
        
        
        graph = {
        student_id: MaxHeap([])
        }
        """
        graph = defaultdict(MaxHeap)
        for item in items:
            graph[item[0]].push((item[1], item[1]))
        
        output = []
        for student_id in graph:
            agg = graph[student_id].get_top_five_agg()
            output.append([student_id, agg])
        
        return sorted(output, key=lambda x:x[0])