class Solution:
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if not root:
            return root
        
        queue = deque([root])
        
        
        while queue:
            size = len(queue)
            prev = None
            for i in range(size):
                curr = queue.popleft()
                if prev:
                    prev.next = curr
                if curr.left: queue.append(curr.left)
                
                if curr.right: queue.append(curr.right)
                prev = curr
            prev = None
        return root