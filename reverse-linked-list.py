# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def recursive_reversal(self, prev, curr):

        while curr:
            temp = curr.next
            curr.next = prev
            prev = curr
            curr = prev

        
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        
        prev = None
        curr = head
        
        while curr:
            temp = curr.next
            curr.next = prev
            # swap that order bitcch
            prev = curr
            curr = temp
        return prev
    
