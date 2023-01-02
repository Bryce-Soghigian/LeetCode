# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: ListNode, left: int, right: int) -> ListNode:
        stack = []
        fast = head
        pt = 1
        while pt != right+1:
            if pt >= left:
                stack.append(fast.val)
            fast = fast.next
            pt += 1
        
        curr = head
        pt = 1
        while pt != right + 1: 
            if pt >=left:
                curr.val = stack.pop()
            curr = curr.next
            pt += 1
        return head