"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        if node == None:
            return node
        graph = {None:None,}
        
        
        # Create all of the nodes
        
        seen = set()
        
        queue = deque([node])
        while queue:
            curr = queue.popleft()
            seen.add(curr.val)
            graph[curr.val] = Node(val=curr.val, neighbors = set())
            for child in curr.neighbors:
                if child.val not in seen:
                    queue.append(child)
        # for each node we want to add teh children
        seen = set()
        
        queue = deque([node])
        while queue:
            curr = queue.popleft()
            seen.add(curr.val)
            # for the current node we want to connect all of the children
            for child in curr.neighbors:
                if child.val not in graph[curr.val].neighbors:
                    graph[curr.val].neighbors.add(graph[child.val])
                
                if child.val not in seen:
                    queue.append(child)
        return graph[node.val]
        