"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        """
        Questions about the question
        Are we going to have duplicate values in our list? 
        Can we Modify the original list?
        
        
        Create a hashmap
        
        This hashmap will map all of our old nodes to the new copy nodes
        
        
        {original_node: copy_node}
        
        We then want to go through and assign the next pointer and random pointer for all of our random nodes
        
        """
        origin_to_new = {None: None}
        
        curr = head
        while curr:
            origin_to_new[curr] = Node(curr.val)
            curr = curr.next
        
        # go through second pass and attach our pointers
        curr = head
        
        while curr:
            # connext new ndoe and random
            
            origin_to_new[curr].next = origin_to_new[curr.next]
            origin_to_new[curr].random = origin_to_new[curr.random]
            curr = curr.next
        
        return origin_to_new[head]
        