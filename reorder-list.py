# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: ListNode) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        
        queue = deque()
        curr = head
        while curr:
            queue.append(curr.val)
            curr = curr.next
        
        curr = head
        bit = 0
        while curr:
            if bit == 0:
                curr.val = queue.popleft()
            else:
                curr.val = queue.pop()
            
            bit ^= 1
            curr = curr.next
        
            
        