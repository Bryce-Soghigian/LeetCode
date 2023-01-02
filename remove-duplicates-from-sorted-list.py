# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if head == None:
            return head
        if head.next == None:
            return head
        """
        We solve this by keeping two pointers
        one which represents next and one that is curr
        if curr == next:
        delete
        else:
        continue
        """
        curr = head
        while curr and curr.next:
            if curr.next.val == curr.val:
                curr.next = curr.next.next
            else:
                curr = curr.next
        return head
        