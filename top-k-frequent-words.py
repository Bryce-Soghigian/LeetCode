class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        """
        Easy solution is to build a max_heap
        """
        output = []
        max_heap = []
        heapify(max_heap)
        count = Counter(words)
        
        for key,value in count.items():
            heappush(max_heap, (-value, key))
        
        
        while k != 0 and max_heap:
            
            curr = heappop(max_heap)
            value, key = curr
            curr_list = [curr]
            while max_heap:
                new = heappop(max_heap)
                if new[0] != value:
                    heappush(max_heap, new)
                    break
                else:
                    curr_list.append(new)
            curr_list.sort(key=lambda x: x[1])
            for i in range(len(curr_list)):
                if k <= 0:
                    break
                output.append(curr_list[i][1])
                k -= 1
        return output
                