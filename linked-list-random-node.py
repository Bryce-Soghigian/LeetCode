"""
Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.

Implement the Solution class:

Solution(ListNode head) Initializes the object with the integer array nums.
int getRandom() Chooses a node randomly from the list and returns its value. All the nodes of the list should be equally likely to be choosen.


"""

import random
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:

    def __init__(self, head):
        self.head = head
        self.nodes = self._build_nodes()
    
    def _build_nodes(self):
        curr = self.head
        output = []
        while curr:
            output.append(curr.val)
            curr = curr.next
        return output
        

    def getRandom(self) -> int:
        return random.choice(self.nodes)