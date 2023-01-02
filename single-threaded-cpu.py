class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        
        time_now = 1
        
        todo = []
        for i, task in enumerate(tasks):
            todo.append((i, task[0], task[1]))
        todo = sorted(todo, key=lambda x: x[1], reverse=True)
        
        queue = []
        order = []
        while todo or queue:
            while todo and time_now >= todo[-1][1]:
                idx, start, process = todo.pop()
                heapq.heappush(queue, (process, idx))
            if queue:
                process, idx = heapq.heappop(queue)
                order.append(idx)
                time_now += process
            else:
                time_now = todo[-1][1]
        
        return order