# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def traverse(self,prev, curr):
        if not curr or not curr.next:
            return
        
        nxt_pair = curr.next.next
        nxt = curr.next
        
        nxt.next = curr
        curr.next = nxt_pair
        prev.next = nxt
        
        
        prev = curr
        curr = nxt_pair
        
        
        

        

        self.traverse(prev, curr)
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        
        sentinel = ListNode(val=None, next=head)
        curr = sentinel.next
        
        self.traverse(sentinel, curr)
        return sentinel.next
        