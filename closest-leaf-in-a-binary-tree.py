class Solution:
    def findClosestLeaf(self, root: TreeNode, k: int) -> int:
        if not root.left and not root.right: return root.val if root.val == k else -1
        adj_list = defaultdict(list)
        stack = [root]
        k_node = None
        while stack:
            node = stack.pop()
            if node.val == k: k_node = node
            if node.left:
                adj_list[node.val].append(node.left)
                adj_list[node.left.val].append(node)
                stack.append(node.left)
            if node.right:
                adj_list[node.val].append(node.right)
                adj_list[node.right.val].append(node)
                stack.append(node.right)
        if not k_node.left and not k_node.right: return k
        q = deque([k_node])

        visited = set([k])
        while q:
            node = q.popleft()
            adj_nodes = adj_list[node.val]
            for adj in adj_nodes:
                if adj.val not in visited:
                    if not adj.left and not adj.right: return adj.val
                    visited.add(adj.val)
                    q.append(adj)
        return -1
               