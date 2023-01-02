def find_node_and_parent(root, key):
    parent = None
    node = root
    while node is not None:
        if node.val < key:
            parent = node
            node = node.right
        elif node.val > key:
            parent = node
            node = node.left
        else:
            return node, parent
    return None, None

def delete_node(node, parent):
    new_node = node.left or node.right
    
    if node.left is not None and node.right is not None:
        rightmost_child = node.left
        while rightmost_child.right is not None:
            rightmost_child = rightmost_child.right
        rightmost_child.right = node.right
    
    return new_node
        
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        dummy = TreeNode(-inf, None, root)
        node, parent = find_node_and_parent(dummy, key)
        
        if node is None:
            return root
        
        new_node = delete_node(node, parent)
        
        if node is parent.left:
            parent.left = new_node
        else:
            parent.right = new_node
        
        return dummy.right