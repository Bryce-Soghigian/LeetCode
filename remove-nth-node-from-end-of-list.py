# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        # get length of list

        length = 1
        curr = head
        prev = head
 
        while curr.next:
            length += 1
            curr = curr.next
            if length > n + 1:
                prev = prev.next
        if length == n:
            return head.next
        else:
            prev.next = prev.next.next
            return head
            